#!/usr/bin/env node
import 'zx/globals';
import { readFile, mkdir } from 'node:fs/promises';
import * as path from 'node:path';

async function main() {
  const fullDigest = process.argv[2];
  const ociLayoutPath = './src';
  const unpackDir = './.unpack';

  if (!fullDigest) {
    console.error('Usage: copy-files-from-image.mts <digest>');
    process.exitCode = 1;
    return;
  }

  const { algorithm, digest } = parseDigest(fullDigest);
  const blobPath = path.join(ociLayoutPath, 'blobs', algorithm, digest);

  if (algorithm.trim() === '' || digest.trim() === '') {
    throw new Error(
      `Invalid digest! ${JSON.stringify({ algorithm, digest })}.}`,
    );
  }

  console.info(`Reading blob file at ${blobPath}`);

  const rawBlob = await readFile(blobPath, { encoding: 'utf8' });
  const jsonBlob = JSON.parse(rawBlob);

  console.dir(jsonBlob, { depth: Infinity });

  const unpackPath = path.join(unpackDir, digest);
  await mkdir(unpackPath, { recursive: true });
  console.info(`unpacking image manifest at ${unpackPath}`);
  // todo: $.sync`regctl image export ocidir://./src@sha256:53ae7d4830054155f58961e01b7a21d182909bab553a9e03af0543e2d08cb644 - | tar -xpf - opt/bitnami/kubectl`
}

function parseDigest(fullDigest: string) {
  const [algorithm, digest] = fullDigest.split(':');

  return {
    algorithm: algorithm ?? '',
    digest: digest ?? '',
  };
}

main().then();
