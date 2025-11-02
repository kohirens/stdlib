# fsio

Package with generic file IO functions. The package was named `fsio` to reduce
conflict with other package names like file, filepath, and other generic names.

## Functions

**Sha256(filePath string) ([]byte, error)** - 
Sha256 returns a SHA256 hash for a file as a byte array. It will throw errors
when it cannot open the file or failing to write the file contents to the hash.
