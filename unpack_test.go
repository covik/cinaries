package main

import (
	"context"
	"os"
	"path/filepath"
	"testing"

	"github.com/opencontainers/go-digest"
	ispec "github.com/opencontainers/image-spec/specs-go/v1"
	rspec "github.com/opencontainers/runtime-spec/specs-go"

	"github.com/opencontainers/umoci/oci/cas/dir"
	"github.com/opencontainers/umoci/oci/casext"
	"github.com/opencontainers/umoci/oci/layer"
)

func TestUnpackRootfs_FromTestdataKubectl(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	testDataPath := filepath.Join("testdata")
	imagePath := filepath.Join(testDataPath, "kubectl-src")

	engine, err := dir.Open(imagePath)
	if err != nil {
		t.Fatalf("open oci layout %q: %v", imagePath, err)
	}
	engineExt := casext.NewEngine(engine)

	manifestDigest := digest.Digest("sha256:53ae7d4830054155f58961e01b7a21d182909bab553a9e03af0543e2d08cb644")
	if err := manifestDigest.Validate(); err != nil {
		t.Fatalf("invalid manifest digest: %v", err)
	}

	// Determine blob size from layout to build a correct descriptor.
	blobPath := filepath.Join(imagePath, "blobs", manifestDigest.Algorithm().String(), manifestDigest.Encoded())
	st, err := os.Stat(blobPath)
	if err != nil {
		t.Fatalf("stat manifest blob: %v", err)
	}
	desc := ispec.Descriptor{
		MediaType: ispec.MediaTypeImageManifest,
		Digest:    manifestDigest,
		Size:      st.Size(),
	}

	blob, err := engineExt.FromDescriptor(ctx, desc)
	if err != nil {
		t.Fatalf("load manifest: %v", err)
	}
	defer func() { _ = blob.Close() }()

	manifest, ok := blob.Data.(ispec.Manifest)
	if !ok {
		t.Fatalf("unexpected blob type for manifest: %T", blob.Data)
	}
	if len(manifest.Layers) == 0 {
		t.Fatalf("manifest has no layers")
	}

	rootfs := filepath.Join(testDataPath, "kubectl-rootfs", manifestDigest.Encoded())
	err = os.MkdirAll(rootfs, os.ModePerm)
	if err != nil {
		t.Fatalf("cannot create dir for rootfs: %v", err)
	}

	unpackOpts := &layer.UnpackOptions{
		OnDiskFormat: layer.DirRootfs{
			MapOptions: layer.MapOptions{
				UIDMappings: []rspec.LinuxIDMapping{{HostID: uint32(os.Geteuid()), ContainerID: 0, Size: 1}},
				GIDMappings: []rspec.LinuxIDMapping{{HostID: uint32(os.Getegid()), ContainerID: 0, Size: 1}},
				Rootless:    os.Geteuid() != 0,
			},
		},
	}

	if err := layer.UnpackRootfs(ctx, engineExt, rootfs, manifest, unpackOpts); err != nil {
		t.Fatalf("UnpackRootfs failed: %v", err)
	}

	// Sanity-check that we unpacked some content.
	entries, err := os.ReadDir(rootfs)
	if err != nil {
		t.Fatalf("read unpacked rootfs: %v", err)
	}
	if len(entries) == 0 {
		t.Fatalf("unpacked rootfs is empty")
	}
}
