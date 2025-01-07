# Setting up development environment

1. Requires [Docker](https://docs.docker.com/engine/install/)
2. Requires Go(lang) from OS package manager or from [tar ball](https://go.dev/doc/install)
3. Install continuous development tool [reflex](https://github.com/cespare/reflex?tab=readme-ov-file#installation)
4. Install [pnpm](https://pnpm.io/installation#using-a-standalone-script) (version 10.0.0)
5. Requires NodeJS v22. e.g. `pnpm` can be used to manage node installations:
   ```bash
   pnpm env list --remote 22
   pnpm env add --global 22
   pnpm env use --global 22
   ```
