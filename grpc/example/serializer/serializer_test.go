package serializer

import (
	"monitoria/grpc/example/pb"
	"monitoria/grpc/example/sample"
	"testing"

	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/proto"
)

func TestFileSerializer(t *testing.T) {
	t.Parallel()

	binayFile := "../tmp/laptop.bin"
	jsonFile := "../tmp/laptop.json"

	laptop1 := sample.NewLaptop()

	err := WriteProtobufToBinaryFile(laptop1, binayFile)
	require.NoError(t, err)

	laptop2 := &pb.Laptop{}
	err = ReadProtobufFromBinaryFile(laptop2, binayFile)
	require.NoError(t, err)
	require.True(t, proto.Equal(laptop1, laptop2))

	err = WriteProtobufToJSONFile(laptop1, jsonFile)
	require.NoError(t, err)
}
