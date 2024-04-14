package keeper

import (
	"codenet/x/codenet/types"
	"encoding/binary"

	"cosmossdk.io/store/prefix"
	"github.com/cosmos/cosmos-sdk/runtime"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// AppendPost appends the given encoded data to the store and returns its ID.
// It increments the count of encoded data and sets the ID of the data before storing it.
func (k Keeper) AppendEncodedData(ctx sdk.Context, data types.EncodedData) uint64 {
	// Get the current count of encoded data
	count := k.GetEncodedDataCount(ctx)
	// Set the ID of the data
	data.DataId = count
	// Create a new store adapter for the provided context
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	// Create a new store with a prefix for encoded data keys
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.EncodeDataKey))
	// Marshal the encoded data into bytes
	appendedValue := k.cdc.MustMarshal(&data)
	// Set the data in the store with its ID as the key
	store.Set(GetEncodedDataIDBytes(data.DataId), appendedValue)
	// Increment the count of encoded data
	k.SeEncodedDataCount(ctx, count+1)
	// Return the ID of the appended data
	return count
}

// GetEncodedDataCount retrieves the current count of encoded data from the store.
func (k Keeper) GetEncodedDataCount(ctx sdk.Context) uint64 {
	// Create a new store adapter for the provided context
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	// Create a new store without any prefix
	store := prefix.NewStore(storeAdapter, []byte{})
	// Retrieve the byte slice associated with the encoded data count key
	byteKey := types.KeyPrefix(types.EncodeDataCountKey)
	bz := store.Get(byteKey)
	// If the byte slice is nil, return 0
	if bz == nil {
		return 0
	}
	// Unmarshal the byte slice into a uint64 and return it
	return binary.BigEndian.Uint64(bz)
}

// GetEncodedDataIDBytes converts a uint64 ID into a byte slice.
func GetEncodedDataIDBytes(id uint64) []byte {
	// Create a new byte slice of length 8
	bz := make([]byte, 8)
	// Convert the uint64 ID to big-endian byte order and store it in the byte slice
	binary.BigEndian.PutUint64(bz, id)
	// Return the byte slice
	return bz
}

// SeEncodedDataCount sets the count of encoded data in the store.
func (k Keeper) SeEncodedDataCount(ctx sdk.Context, count uint64) {
	// Create a new store adapter for the provided context
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	// Create a new store without any prefix
	store := prefix.NewStore(storeAdapter, []byte{})
	// Retrieve the byte slice associated with the encoded data count key
	byteKey := types.KeyPrefix(types.EncodeDataCountKey)
	// Create a new byte slice of length 8 and store the count in big-endian byte order
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	// Set the byte slice in the store
	store.Set(byteKey, bz)
}
