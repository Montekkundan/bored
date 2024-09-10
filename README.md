# Bored


## Available Commands

Turborepo allows you to run commands across all projects or filter for specific projects. Here are the main commands available:

- `turbo dev`: Start both web and backend in development mode
- `turbo build`: Build the web project
- `turbo lint`: Run linting for the web project
- `turbo start`: Start both web and backend applications
- `turbo stop --filter=backend`: Stop the backend application

## Usage Examples

### Running All Projects

To start both the web and backend in development mode:

`turbo dev`

### Filtering Commands

You can use the `--filter` flag to run commands for specific projects:

- Start only the web project in development mode:

`turbo dev --filter=web`

- Start only the backend in development mode:

`turbo dev --filter=backend`