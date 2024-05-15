import crypto from 'crypto';

export function generateSHA256Hash() {
    // Create a hash buffer:
    const randomBuffer = crypto.randomBytes(1024);
    // Create hash using sha256:
    const hash = crypto.createHash('sha256');
    hash.update(randomBuffer);
    const hexHash = hash.digest('hex');

    return hexHash;
}