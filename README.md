# mod-cleaner

Clean go modules

## Inputs

### `gomods`

**Optional** glob pattern to match `go.mod` files to examine.

## Example usage

uses: actions/mod-cleaner@v1
with:
  gomods: 'go.mod'
