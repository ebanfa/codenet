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
	count := k.GetCountOfEncodedData(ctx)
	// Set the ID of the data
	data.EncodedDataId = count
	// Create a new store adapter for the provided context
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	// Create a new store with a prefix for encoded data keys
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.EncodedDataKey))
	// Marshal the encoded data into bytes
	appendedValue := k.cdc.MustMarshal(&data)
	// Set the data in the store with its ID as the key
	store.Set(GetEncodedDataIDBytes(data.EncodedDataId), appendedValue)
	// Increment the count of encoded data
	k.SeEncodedDataCount(ctx, count+1)
	// Return the ID of the appended data
	return count
}

// GetCountOfEncodedData retrieves the current count of encoded data from the store.
func (k Keeper) GetCountOfEncodedData(ctx sdk.Context) uint64 {
	// Create a new store adapter for the provided context
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	// Create a new store without any prefix
	store := prefix.NewStore(storeAdapter, []byte{})
	// Retrieve the byte slice associated with the encoded data count key
	byteKey := types.KeyPrefix(types.EncodedDataCountKey)
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
	byteKey := types.KeyPrefix(types.EncodedDataCountKey)
	// Create a new byte slice of length 8 and store the count in big-endian byte order
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	// Set the byte slice in the store
	store.Set(byteKey, bz)
}

// GetPost retrieves the encoded data with the given ID from the store.
func (k Keeper) GetEncodedData(ctx sdk.Context, id uint64) (val types.EncodedData, found bool) {
	// Create a new store adapter for the provided context
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	// Create a new store with a prefix for encoded data keys
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.EncodedDataKey))
	// Retrieve the byte slice associated with the encoded data ID key
	b := store.Get(GetEncodedDataIDBytes(id))
	// If the byte slice is nil, return an empty EncodedData and false
	if b == nil {
		return val, false
	}
	// Unmarshal the byte slice into EncodedData and return it along with true
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// SetEncodedData stores the provided encoded data in the store.
func (k Keeper) SetEncodedData(ctx sdk.Context, encodedData types.EncodedData) {
	// Create a new store adapter for the provided context
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	// Create a new store with a prefix for encoded data keys
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.EncodedDataKey))
	// Marshal the encoded data into bytes
	b := k.cdc.MustMarshal(&encodedData)
	// Set the encoded data in the store with its ID as the key
	store.Set(GetEncodedDataIDBytes(encodedData.EncodedDataId), b)
}
