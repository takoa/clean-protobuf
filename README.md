# Clean-Protobuf

A clean architecture version of https://github.com/grpc/grpc-go/tree/master/examples/route_guide.

The project layout is based on https://github.com/golang-standards/project-layout.

The architecture is influenced by https://github.com/evrone/go-clean-template, and a little bit https://github.com/Creatly/creatly-backend.

## Layout
### `/api`
The API definition file (like `.proto`), and its generated code.

### `/assets`
Non-Go related data. Usually things like images, but here it is test data.

### `/cmd`
The mains.

### `/internal`
The implementation not meant to be used by external systems.

#### `/internal/app`
The starting points of the systems. Basically extensions of the mains.

#### `/internal/config`
The global configurations.

#### `/internal/entity`
The entities of the system. Sometimes named as domain.

The package contains definitions of the common types and their methods.

#### `/internal/entity/model`
The models. Methods for manipulating models also resides here.

#### `/internal/entity/repository`
The repository interfaces.

Actual implementation requires communication with the outside world (a database, etc) thus resides in `/internal/infrastructure/repository`.

#### `/internal/infrastructure`
The implementation which directly communicates with the outside world and does conversion into entities.

Basically gateways/mediators between the business logic (use cases) and the outside world that translates "their" data into "ours".

#### `/internal/infrastructure/controller`
The controllers/handlers of the server.

#### `/internal/infrastructure/repository`
The database logic. It reads/writes data from/to the databases or alike.

#### `/internal/pkg`
The internal packages.

The packages are independent of the rest of the implementation.

#### `/internal/usecase`
The business logic. The part which is not "chores".

It does not have direct external dependencies like database or API definitions. All those dependencies must be abstracted to be used here.

## Notes
### Use Case Callbacks
Use cases call a callback function when an important event occurs. Callbacks act as output ports, abstraction of the outside world, and thus are supposed to be implemented and provided by the infrastructure layer.

This style is rather rigorous and cumbersome. Moreover, majority of the times, it can be replaced by simple return values. However, it provides a lot more flexibility to the infrastructure layer. The examples are shown in the streaming RPC controllers where we still achieve practically one-to-one port from the original single handler implementation to infrastructure-usecase-entity implementation.

### Differences from the Original
- `RouteChat` (`route.PostMessage`) locks its underlying map separately when saving a new message and reading existing messages. The original locks both together.
    - While the behavior is slightly different, there's no practical disadvantage (your new message might not be the latest message, but who cares?)
- Use environment variables instead of flags.
    - Having a struct felt cleaner while providing `/internal/config` example.
- General refactoring/renaming