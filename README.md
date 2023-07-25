# sabal-pb
gRPC library

# EquipService Definition

This README provides an overview of the `EquipService` ProtoBuf definition used in the `pb` package of this application.

## File: EquipService.proto

The EquipService defines a service called `EquipService`, which includes several Remote Procedure Call (RPC) methods for interacting with `Equip` entities.

The RPC methods defined in the service are:

1. **Get**: Accepts an `EquipId` and returns an `Equip` object. This method is used to fetch an equipment entity based on its ID.

2. **GetAll**: Accepts a `ListRequest` object and returns an `EquipList`. This method is used to fetch a list of all equipment entities.

3. **Create**: Accepts an `Equip` object and returns an `Equip` object. This method is used to create a new equipment entity.

4. **Update**: Accepts an `Equip` object and returns an `Equip` object. This method is used to update an existing equipment entity.

5. **Delete**: Accepts an `EquipId` and returns a `DeleteResponse`. This method is used to delete an equipment entity based on its ID.

6. **GetOrCreate**: Accepts an `Equip` object and returns an `Equip` object. This method is used to fetch an equipment entity if it exists or create a new one if it does not.

Note: The actual ProtoBuf message definitions for `Equip`, `EquipId`, `ListRequest`, `EquipList`, and `DeleteResponse` are located in the imported `common.proto` and `equip.proto` files. Make sure to refer to these files for more information on the message structures.
