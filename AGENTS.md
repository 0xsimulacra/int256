# AGENTS.md

## Global Codex workflow

When working in any software project, always detect and use the project-provided development environment before running tools, tests, linters, formatters, builds, language servers, package managers, or scripts.

Do not assume globally installed tools are the correct tools for the project.

## Working directory

Before running project commands, move to the project root.

Prefer the Git repository root when available:

```sh
cd "$(git rev-parse --show-toplevel)"
```

If the current directory is not inside a Git repository, stay in the current project directory and explain that no Git root was found.

## Operating system gate

Before using Nix-specific project setup, first check whether the current system is actually NixOS.

Treat the system as NixOS if one of these checks succeeds:

```sh
test -e /etc/NIXOS
grep -q '^ID=nixos$' /etc/os-release
```

If the system is not NixOS, do not run `direnv`, `nix develop`, `nix build`, or other Nix-specific setup commands unless the user explicitly asks for them.

On non-NixOS systems, ignore `.envrc` and `flake.nix` for automatic setup and instead follow the project’s documented tooling, such as README instructions, package scripts, Makefile targets, language-specific tooling, or commands explicitly provided by the user.

## Environment loading priority

After moving to the project root, choose the environment setup using this priority:

1. If the system is NixOS and the project has a `.envrc`, use `direnv`.
2. If the system is NixOS and the project has no `.envrc` but has a `flake.nix`, use `nix develop`.
3. Otherwise, follow the project’s documented tooling.

Do not use `direnv` or `nix develop` automatically on non-NixOS systems.

## Projects with `.envrc` on NixOS

If the system is NixOS and a `.envrc` file exists at the project root, use `direnv`, but do not run `direnv reload` by default.

For normal one-off commands, prefer:

```sh
direnv exec . <command>
```

Examples:

```sh
direnv exec . go test ./...
direnv exec . cargo test
direnv exec . bun test
direnv exec . bun run lint
direnv exec . forge test
```

Only run:

```sh
direnv allow
```

when direnv reports that the `.envrc` is blocked or not allowed.

Only run:

```sh
direnv reload
```

when one of these is true:

* `.envrc` was changed
* `direnv allow` was just run and the environment still needs loading
* `direnv exec . <command>` fails because the environment is stale
* the user explicitly asks to reload the environment

Do not run `direnv reload` as routine setup before every command.

Do not use `direnv reload` to “prepare” the user’s interactive shell. The agent should run commands through `direnv exec .` instead.

Do not use `.direnv/` as a source directory. It is cache/build environment state, not project source.

Do not edit `.envrc`, `.env`, `flake.nix`, or `flake.lock` unless the task explicitly requires it.

Do not run commands that update the Nix flake lock file unless explicitly asked, such as:

```sh
nix flake update
nix flake lock
```

Avoid `--refresh` for normal checks unless explicitly asked.

If `.envrc` uses `use flake`, remember that this loads a build environment similar to `nix develop`, so do not also run `nix develop` manually for the same command path. Use `direnv exec . <command>`.


## Projects without `.envrc` but with `flake.nix` on NixOS

If the system is NixOS, there is no `.envrc`, and there is a `flake.nix`, use the Nix flake development shell.

For one-off commands, prefer:

```sh
nix develop . --command <command>
```

Examples:

```sh
nix develop . --command go test ./...
nix develop . --command cargo test
nix develop . --command bun test
nix develop . --command bun run lint
nix develop . --command forge test
```

For an interactive shell, use:

```sh
nix develop .
```

Do not run `nix build` unless explicitly asked.

## Non-NixOS projects

If the system is not NixOS, do not automatically use `direnv` or `nix develop`, even if `.envrc` or `flake.nix` exists.

Instead, inspect and follow the project’s documented tooling in this order when available:

1. User instructions in the current task.
2. `AGENTS.md` or other repo-specific agent instructions.
3. `README.md`, `CONTRIBUTING.md`, or developer docs.
4. `Makefile`, `justfile`, package scripts, or language-specific commands.
5. The smallest reasonable standard command for the detected language or framework.

Examples of non-NixOS fallback commands may include:

```sh
go test ./...
cargo test
bun test
bun run lint
forge test
make test
just test
```

Only use these after checking the project’s own documentation or scripts.

## Tooling preferences

Prefer project-local tools over global tools when possible.

On NixOS, prefer flake-provided tools through `direnv` or `nix develop`.

On non-NixOS, prefer the tooling documented by the project.

Prefer:

```sh
bun
```

over:

```sh
npm
```

when the project supports Bun.

Use `bun install`, `bun test`, and `bun run ...` when applicable.


## Command pattern

Recommended pattern for one-off commands on NixOS with `.envrc`:

```sh
cd "$(git rev-parse --show-toplevel)"
direnv exec . <command>
```

If direnv says the `.envrc` is not allowed:

```sh
cd "$(git rev-parse --show-toplevel)"
direnv allow
direnv exec . <command>
```

If `direnv exec` still reports a stale or failed environment after allowing:

```sh
cd "$(git rev-parse --show-toplevel)"
direnv reload
direnv exec . <command>
```

Recommended pattern for one-off commands on NixOS without `.envrc` but with `flake.nix`:

```sh
cd "$(git rev-parse --show-toplevel)"
nix develop . --command <command>
```

Recommended setup pattern on non-NixOS:

```sh
cd "$(git rev-parse --show-toplevel)"
# Do not run direnv or nix develop automatically.
# Follow the project documentation and scripts instead.
```

## Command style

Avoid hiding errors with shell patterns such as:

```sh
2>/dev/null || true
```

Prefer commands that fail clearly and show useful output.

Do not suppress errors unless explicitly asked.

## Code editing expectations

* Prefer small, direct changes.
* Avoid unnecessary rewrites.
* Preserve existing project conventions.
* Keep existing function names stable unless a rename is explicitly requested.
* Do not rename functions that do the same thing.
* Avoid adding broad defensive validation unless it is needed or requested.
* Prefer Tailwind CSS utilities over bare CSS in frontend UI.
* Prefer Bun over npm in JavaScript and TypeScript projects.

## Testing expectations

After changing code, run the smallest relevant test first.

On NixOS with `.envrc`, examples:

```sh
direnv exec . go test ./path/to/package
direnv exec . cargo test -p crate_name
direnv exec . bun test
direnv exec . forge test --match-test SomeTest
```

On NixOS with `flake.nix` but no `.envrc`, examples:

```sh
nix develop . --command go test ./path/to/package
nix develop . --command cargo test -p crate_name
nix develop . --command bun test
nix develop . --command forge test --match-test SomeTest
```

On non-NixOS, use the project’s documented test command. Examples only when appropriate:

```sh
go test ./path/to/package
cargo test -p crate_name
bun test
forge test --match-test SomeTest
make test
just test
```

If the small test passes and the change is broader, run the wider relevant suite.

On NixOS with `.envrc`, examples:

```sh
direnv exec . go test ./...
direnv exec . cargo test
direnv exec . bun run lint
direnv exec . forge test
```

On NixOS with `flake.nix` but no `.envrc`, examples:

```sh
nix develop . --command go test ./...
nix develop . --command cargo test
nix develop . --command bun run lint
nix develop . --command forge test
```

On non-NixOS, follow the project documentation or scripts.

## Summary rule

Always start from the project root.

First check whether the system is NixOS.

If the system is NixOS and `.envrc` exists, run commands with:

```sh
direnv exec . <command>
```

If direnv says the `.envrc` is not allowed:

```sh
direnv allow
direnv exec . <command>
```

Only use `direnv reload` when explicitly needed.

If the system is NixOS, `.envrc` does not exist, and `flake.nix` exists, run commands with:

```sh
nix develop . --command <command>
```

If the system is not NixOS, do not automatically run `direnv` or `nix develop`. Follow the project’s documented tooling instead.


## TypeScript / Bun projects

If a project uses TypeScript and Bun, prefer Bun commands over npm commands.

Prefer:

```sh
bun install
bun test
bun run lint
bun run build
bun run typecheck
```

over npm equivalents, unless the project documentation explicitly says otherwise.

### Dependency source inspection for TypeScript / Bun projects

`node_modules/` may be available for targeted inspection, but it must not be scanned broadly by default.

Default behavior:

```txt
Do not scan node_modules by default.
Do not grep the whole node_modules tree.
Only inspect a specific package inside node_modules when dependency behavior, types, exports, or source code are directly relevant.
```

Reading the installed package version can be useful for JS/TS bugs, especially when docs, GitHub, npm, or examples do not match what is actually installed. But freely searching all of `node_modules/` creates noise, slows down work, and burns context.

When searching the project, exclude dependency/vendor directories unless there is a specific reason to inspect them.

Prefer searches like:

```sh
rg "pattern" . --glob '!node_modules/**' --glob '!vendor/**' --glob '!dist/**' --glob '!build/**'
```

Only read files under `node_modules/` when the task specifically depends on the installed dependency implementation, types, exports, package metadata, or generated code.

Good reasons to inspect `node_modules/` include:

* checking the exact installed package version behavior
* reading `.d.ts` files for confusing TypeScript errors
* checking package `exports` / `types` / `module` / `main`
* debugging bundler resolution
* verifying whether docs differ from installed code
* inspecting patched dependencies

When inspecting `node_modules/`, target the smallest relevant path.

Prefer:

```sh
cat node_modules/<package>/package.json
rg "symbolName" node_modules/<package> --glob '!**/*.map'
```

Avoid:

```sh
rg "pattern" node_modules
find node_modules
ls -R node_modules
```

Never edit files under `node_modules/`.

If a dependency needs to be changed, modify the real source, apply a proper patch mechanism, use package-manager patching, or change the project configuration instead.


Add a section like this to `AGENTS.md`:


## Dependency source inspection for Go / Rust / Foundry projects

For Go, Rust, and Foundry projects, dependency source may be inspected when it is directly useful, but it must be done moderately and through language-aware tooling.

Do not broadly scan dependency caches, build outputs, `.direnv/`, Nix store paths, or generated directories.

Default behavior:

```txt
Do not scan dependency caches by default.
Do not grep the whole Go module cache, Cargo registry, Foundry lib directory, .direnv, target, out, or cache trees.
Only inspect a specific dependency/package/crate/contract source when behavior, types, ABI, interfaces, or implementation details are directly relevant.
```

The project may use `direnv`, so dependency locations may come from the development environment. Use the project environment first, then ask the language tooling where the source is.

Do not manually inspect `.direnv/`. It is an environment/cache directory, not project source.

Do not inspect Nix store paths unless the task specifically requires debugging packaging, toolchain, or generated dependency source behavior.

Never edit dependency cache files.

If a dependency needs to be changed, modify the real project source, use the project’s patch mechanism, fork/override the dependency, or change configuration.

### Go projects

Use Go tooling to find dependency source instead of guessing paths.

Prefer:

```sh
go env GOPATH GOMODCACHE GOROOT
go list -m -f '{{.Path}} {{.Version}} {{.Dir}}' all
go list -f '{{.ImportPath}} {{.Dir}}' ./...
go doc <package-or-symbol>
```

For a specific dependency, prefer targeted inspection:

```sh
go list -m -f '{{.Dir}}' <module>
rg "symbolName" "$(go list -m -f '{{.Dir}}' <module>)"
```

Good reasons to inspect Go dependency source include:

* checking exact installed module behavior
* reading interfaces, structs, or generated code
* debugging build tags or platform-specific files
* verifying behavior that differs from documentation
* understanding serialization, RPC, ABI, or low-level implementation details

Avoid:

```sh
rg "pattern" "$(go env GOMODCACHE)"
find "$(go env GOMODCACHE)"
ls -R "$(go env GOMODCACHE)"
```

Never edit files in the Go module cache.

Follow the existing code style and naming style in the repository. Do not rename functions, types, variables, files, or concepts unless the user explicitly asks for a rename.

Do not add defensive nil checks, zero-value fallbacks, bounds guards, or error swallowing for values that are not expected to be nil/invalid in normal execution. If such a value is unexpectedly nil or invalid, prefer letting the program fail loudly so the bug is caught immediately instead of hiding the failure and continuing with corrupted assumptions.

Only add nil checks or defensive validation when:

* the value is genuinely optional by API contract
* the existing code already treats it as optional
* the check is required at an external boundary, such as RPC, JSON, CLI input, database input, or untrusted calldata
* the user explicitly asks for defensive handling

For internal hot-path logic, prefer clear invariants, direct code, and fail-fast behavior over extra helper layers or defensive branches.

### Rust projects

Use Cargo and Rust tooling to find dependency source instead of guessing paths.

Prefer:

```sh
cargo metadata --format-version=1
cargo tree
rustc --print sysroot
````

For a specific crate, inspect only the relevant crate source from the Cargo registry, git checkout, workspace, or sysroot path discovered by the tooling.

Good reasons to inspect Rust dependency source include:

* checking exact installed crate behavior
* reading trait definitions, macros, or feature-gated code
* debugging feature flags
* understanding generated code or proc macro behavior
* reading standard library source when relevant

Avoid broad searches like:

```sh
rg "pattern" ~/.cargo
find ~/.cargo
ls -R ~/.cargo
rg "pattern" target
```

Do not scan the whole Cargo registry, git cache, `target/`, or `.direnv/`.

Never edit files in the Cargo registry, Cargo git cache, Rust sysroot, or `target/`.

#### Rust workspace style

`lib.rs` files must be minimal and should not contain meaningful logic.

Use this for the crate-level documentation in `lib.rs`:

```rust
#![doc = include_str!("../README.md")]
```

Do not use `//!` crate-level comments in `lib.rs`.

Group each module declaration with its re-export:

```rust
mod foo;
pub use foo::Bar;

mod baz;
pub use baz::{Baz, BazError};
```

Do not list all `mod` declarations first and all `pub use` statements later.

Prefer private modules with explicit re-exports from `lib.rs`.

Only make modules `pub` or `pub(crate)` when the module namespace itself is intentionally part of the API, or for test utilities such as:

```rust
pub mod test_utils;
```

Public API items should be `pub` and re-exported from `lib.rs` when they are part of the crate API.

Internal helpers may remain private when they are implementation details.

Avoid `pub(crate)` unless there is a clear cross-module need.

Prefer type-centered public APIs. Put functions as methods on a type when that makes the API clearer.

Bare public functions are allowed when they are clearer than introducing an artificial unit struct.

Do not add `#![allow(missing_docs)]` or broad allow-lints to suppress clippy or rustdoc warnings. Fix the underlying issue instead.

Binary crates under `bin/` should contain minimal glue code. All meaningful logic belongs in library crates.

Cargo.toml dependencies should follow the existing workspace style, including line-length waterfall sorting when that style is used.

Logically group dependencies as done in the rest of the workspace.

Features sections go at the bottom of the manifest.

All crate and binary `Cargo.toml` files must inherit lints from the workspace:

```toml
[lints]
workspace = true
```

Do not add dependency features in the workspace root `Cargo.toml`.

Enable dependency features only in the individual crates or binaries that need them, to prevent feature leakage into `no_std` crates.

Every `mod.rs` file must begin with a `//!` module doc comment describing what the module contains.

All `use` imports must be at the top of the file or the top of a `mod` block.

Never place `use` statements inside function bodies or closures.

Exceptions:

* conditional imports inside `#[cfg(...)]` blocks are allowed
* imports inside `#[cfg(test)] mod tests` are allowed
* imports inside feature-gated modules/functions are allowed
* `use` inside `macro_rules!` bodies is allowed when needed by the macro expansion

Use structured tracing instead of interpolated strings.

Always use key-value fields for dynamic data:

```rust
info!(block = %block_number, "processed block");
error!(error = %e, peer = %peer_id, "connection failed");
```

Use `%` for `Display` and `?` for `Debug`.

The tracing message string should be static. Put variable data in fields.

Avoid:

```rust
info!("processed block {block_number}");
error!("connection to {peer_id} failed: {e}");
```

`#[cfg(test)] mod tests { ... }` must be placed at the end of the file, after all non-test code.


### Foundry / Solidity projects

Use Foundry tooling before manually inspecting dependency paths.

The project may use Soldeer for Solidity dependencies. In most cases, dependencies are already fetched under `dependencies/`.

Do not fetch, update, or reinstall Soldeer dependencies unless explicitly asked.

Before adding or changing Solidity imports, always inspect the project remappings first.

Prefer:

```sh
forge remappings
forge tree
forge inspect <ContractName> abi
forge inspect <ContractName> storageLayout
forge inspect <ContractName> methods
```

Use the remapping prefix instead of importing from the raw `dependencies/` path.

Example remappings may look like:

```txt
@blocknumberish/=dependencies/blocknumberish-v0.x.0/
@forge-std/=dependencies/forge-std-1.16.1/src/
@openzeppelin-contracts-480/=dependencies/openzeppelin-contracts-480-4.8.0/contracts/
@openzeppelin-contracts-502/=dependencies/openzeppelin-contracts-502-5.0.2/contracts/
@solady/=dependencies/solady-0.1.26+/
@solmate/=dependencies/solmate-v0.6.8/
@testSrcRoot/=src/test/
@uni-mixed-quoter/=dependencies/uni-mixed-quoter-v0.x.0/
```

If a file exists at:

```txt
dependencies/openzeppelin-contracts-502-5.0.2/contracts/token/ERC20/ERC20.sol
```

and the remapping is:

```txt
@openzeppelin-contracts-502/=dependencies/openzeppelin-contracts-502-5.0.2/contracts/
```

then import it through the remapping:

```solidity
import { ERC20 } from "@openzeppelin-contracts-502/token/ERC20/ERC20.sol";
```

Do not write imports directly against `dependencies/...` unless the project already consistently does that.

Preserve exact file and directory casing when writing Solidity imports.

Read dependency source only when the specific contract, interface, library, or remapping is directly relevant.

Good reasons to inspect Foundry dependency source include:

* checking exact interface or library behavior
* reading inherited contracts
* verifying remappings
* understanding ABI, storage layout, modifiers, or internal library logic
* debugging why a contract compiles or tests differently than expected
* confirming the correct import path before adding an import

Prefer targeted searches:

```sh
rg "contract ContractName" src test script dependencies
rg "interface InterfaceName" src test script dependencies
rg "library LibraryName" src test script dependencies
rg "function functionName" src test script dependencies/<specific-dependency>
```

Avoid broad dependency/cache searches:

```sh
rg "pattern" dependencies
find dependencies
ls -R dependencies
rg "pattern" out
rg "pattern" cache
```

Do not broadly search `lib/` when it is being used as a dependency/vendor directory.

If `lib/` is project-owned source, targeted searches and edits in `lib/` are fine.

Never edit third-party dependency files under `dependencies/`.

Before editing files under `lib/`, determine whether the file is project-owned or third-party dependency code by checking remappings, imports, git status, package metadata, and surrounding project conventions.

If a third-party Solidity dependency needs to change, use a fork, remapping, patch, Soldeer override, or project-level override instead.

Never edit `out/`, `cache/`, or generated artifacts.
