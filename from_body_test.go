package restpc

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/golang/mock/gomock"
)

func TestFromBody_GetString(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockReq := NewMockRequest(ctrl)
	var req Request = mockReq
	{
		mockReq.EXPECT().BodyMap().Return(nil, fmt.Errorf("unknown error"))
		value, err := FromBody.GetString(req, "name")
		assert.Nil(t, value)
		assert.EqualError(t, err, "unknown error")
	}
	{
		mockReq.EXPECT().BodyMap().Return(nil, nil)
		value, err := FromBody.GetString(req, "name")
		assert.Nil(t, value)
		assert.Nil(t, err)
	}
	{
		mockReq.EXPECT().BodyMap().Return(map[string]interface{}{
			"name": 123,
		}, nil)
		value, err := FromBody.GetString(req, "name")
		assert.Nil(t, value)
		assert.EqualError(t, err, "invalid 'name', must be string")
	}
	{
		mockReq.EXPECT().BodyMap().Return(map[string]interface{}{
			"name": "John Smith",
		}, nil)
		value, err := FromBody.GetString(req, "name")
		assert.Equal(t, "John Smith", *value)
		assert.Nil(t, err)
	}
	{
		mockReq.EXPECT().BodyMap().Return(map[string]interface{}{
			"name": []byte("John Smith"),
		}, nil)
		value, err := FromBody.GetString(req, "name")
		assert.Equal(t, "John Smith", *value)
		assert.Nil(t, err)
	}
}
func TestFromBody_GetStringList(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockReq := NewMockRequest(ctrl)
	var req Request = mockReq
	{
		mockReq.EXPECT().BodyMap().Return(nil, fmt.Errorf("unknown error"))
		value, err := FromBody.GetStringList(req, "names")
		assert.Nil(t, value)
		assert.EqualError(t, err, "unknown error")
	}
	{
		mockReq.EXPECT().BodyMap().Return(nil, nil)
		value, err := FromBody.GetStringList(req, "names")
		assert.Nil(t, value)
		assert.Nil(t, err)
	}
	{
		mockReq.EXPECT().BodyMap().Return(map[string]interface{}{
			"names": 123,
		}, nil)
		value, err := FromBody.GetStringList(req, "names")
		assert.Nil(t, value)
		assert.EqualError(t, err, "invalid 'names', must be array of strings")
	}
	{
		mockReq.EXPECT().BodyMap().Return(map[string]interface{}{
			"names": "John Smith",
		}, nil)
		value, err := FromBody.GetStringList(req, "names")
		assert.Nil(t, value)
		assert.EqualError(t, err, "invalid 'names', must be array of strings")
	}
	{
		mockReq.EXPECT().BodyMap().Return(map[string]interface{}{
			"names": []interface{}{"John Smith", 1234},
		}, nil)
		value, err := FromBody.GetStringList(req, "names")
		assert.Nil(t, value)
		assert.EqualError(t, err, "invalid 'names', must be array of strings")
	}
	{
		mockReq.EXPECT().BodyMap().Return(map[string]interface{}{
			"names": []string{"John Smith", "John Doe"},
		}, nil)
		value, err := FromBody.GetStringList(req, "names")
		assert.Equal(t, []string{"John Smith", "John Doe"}, value)
		assert.Nil(t, err)
	}
	{
		mockReq.EXPECT().BodyMap().Return(map[string]interface{}{
			"names": []interface{}{"John Smith", "John Doe"},
		}, nil)
		value, err := FromBody.GetStringList(req, "names")
		assert.Equal(t, []string{"John Smith", "John Doe"}, value)
		assert.Nil(t, err)
	}

}
