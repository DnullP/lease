```mermaid
classDiagram
namespace Client {
    class CenterNodeClient{
        <<interface>>
        RequestData(ctx, in, opts) response
        WriteData(ctx, in, opts) response
    }
    class centerNodeClient{
        cc grpc.ClientConnInterface
        RequestData(ctx, in, opts) response
        WriteData(ctx, in, opts) response
    }
}
    centerNodeClient --|> CenterNodeClient
    note for centerNodeClient "Created by NewCenterNodeClient()"
```

classDiagram
classA --|> classB : Inheritance
classC --* classD : Composition
classE --o classF : Aggregation
classG --> classH : Association
classI -- classJ : Link(Solid)
classK ..> classL : Dependency
classM ..|> classN : Realization
classO .. classP : Link(Dashed)

classDiagram
class Shape
<<interface>> Shape
Shape : noOfVertices
Shape : draw()
