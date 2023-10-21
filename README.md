# Datum Core

This repo will hold the core server / handler for Datum services.

## Development

Datum's core server operates with the following utilities:

1. [ent](https://entgo.io/) - ORM
1. [atlas](https://atlasgo.io/) - Schema generation and migration
1. [gqlgen](https://gqlgen.com/) - Code generation from schema definitions
1. [openfga](https://openfga.dev/) - Authorization 

### Extending the existing server

1. Create a new schema by running the following command, replacing `<object>` with your object:
```bash
go run -mod=mod entgo.io/ent/cmd/ent new --target internal/ent/schema <object> 
```
For example, if you wanted to create a user, organization, and members schema you would run:
```bash
go run -mod=mod entgo.io/ent/cmd/ent new --target internal/ent/schema User Organization Member 
```
1. This will generate a file per schema in `internal/ent/schema`

```bash
tree internal/ent/schema 

internal/ent/schema
└── user.go
└── organization.go
└── member.go
```

1. You will add your fields, edges, annotations, etc to this file for each schema. See the [ent schema def docs](https://entgo.io/docs/schema-def) for more details. 

1. Now you will need to create a `graphql` file per schema that will handle CRUD operations, using the same example this would look like: 
```
tree schema 
schema
├── ent.graphql
└── user.graphql
└── organization.graphql
└── member.graphql
```

To have the files auto generated, use:

```bash
make graph
```

1. Now that all the code is there, test it using the playground:
```
make run-dev
```
Using the default config, you should be able to go to your browser of choice and see the playground: http://localhost:17608/playground or Via curl, `http://localhost:17608/query`

