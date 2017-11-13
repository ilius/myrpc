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

func TestFromBody_GetInt(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockReq := NewMockRequest(ctrl)
	var req Request = mockReq
	{
		mockReq.EXPECT().BodyMap().Return(nil, fmt.Errorf("unknown error"))
		value, err := FromBody.GetInt(req, "count")
		assert.Nil(t, value)
		assert.EqualError(t, err, "unknown error")
	}
	{
		mockReq.EXPECT().BodyMap().Return(nil, nil)
		value, err := FromBody.GetInt(req, "count")
		assert.Nil(t, value)
		assert.Nil(t, err)
	}
	{
		mockReq.EXPECT().BodyMap().Return(map[string]interface{}{
			"count": "abc",
		}, nil)
		value, err := FromBody.GetInt(req, "count")
		assert.Nil(t, value)
		assert.EqualError(t, err, "invalid 'count', must be integer")
	}
	{
		mockReq.EXPECT().BodyMap().Return(map[string]interface{}{
			"count": "345",
		}, nil)
		value, err := FromBody.GetInt(req, "count")
		assert.Nil(t, value)
		assert.EqualError(t, err, "invalid 'count', must be integer")
	}
	{
		mockReq.EXPECT().BodyMap().Return(map[string]interface{}{
			"count": 5001,
		}, nil)
		value, err := FromBody.GetInt(req, "count")
		assert.Equal(t, 5001, *value)
		assert.Nil(t, err)
	}
	{
		mockReq.EXPECT().BodyMap().Return(map[string]interface{}{
			"count": int32(5003),
		}, nil)
		value, err := FromBody.GetInt(req, "count")
		assert.Equal(t, 5003, *value)
		assert.Nil(t, err)
	}
	{
		mockReq.EXPECT().BodyMap().Return(map[string]interface{}{
			"count": int64(6123),
		}, nil)
		value, err := FromBody.GetInt(req, "count")
		assert.Equal(t, 6123, *value)
		assert.Nil(t, err)
	}
	{
		mockReq.EXPECT().BodyMap().Return(map[string]interface{}{
			"count": 14.15,
		}, nil)
		value, err := FromBody.GetInt(req, "count")
		assert.Equal(t, 14, *value)
		assert.Nil(t, err)
	}
}

func TestFromBody_GetFloat(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockReq := NewMockRequest(ctrl)
	var req Request = mockReq
	{
		mockReq.EXPECT().BodyMap().Return(nil, fmt.Errorf("unknown error"))
		value, err := FromBody.GetFloat(req, "weight")
		assert.Nil(t, value)
		assert.EqualError(t, err, "unknown error")
	}
	{
		mockReq.EXPECT().BodyMap().Return(nil, nil)
		value, err := FromBody.GetFloat(req, "weight")
		assert.Nil(t, value)
		assert.Nil(t, err)
	}
	{
		mockReq.EXPECT().BodyMap().Return(map[string]interface{}{
			"weight": "abc",
		}, nil)
		value, err := FromBody.GetFloat(req, "weight")
		assert.Nil(t, value)
		assert.EqualError(t, err, "invalid 'weight', must be float")
	}
	{
		mockReq.EXPECT().BodyMap().Return(map[string]interface{}{
			"weight": "345",
		}, nil)
		value, err := FromBody.GetFloat(req, "weight")
		assert.Nil(t, value)
		assert.EqualError(t, err, "invalid 'weight', must be float")
	}
	{
		mockReq.EXPECT().BodyMap().Return(map[string]interface{}{
			"weight": 1231,
		}, nil)
		value, err := FromBody.GetFloat(req, "weight")
		assert.Equal(t, 1231.0, *value)
		assert.Nil(t, err)
	}
	{
		mockReq.EXPECT().BodyMap().Return(map[string]interface{}{
			"weight": int32(2345),
		}, nil)
		value, err := FromBody.GetFloat(req, "weight")
		assert.Equal(t, 2345.0, *value)
		assert.Nil(t, err)
	}
	{
		mockReq.EXPECT().BodyMap().Return(map[string]interface{}{
			"weight": int64(7123),
		}, nil)
		value, err := FromBody.GetFloat(req, "weight")
		assert.Equal(t, 7123.0, *value)
		assert.Nil(t, err)
	}
	{
		mockReq.EXPECT().BodyMap().Return(map[string]interface{}{
			"weight": 104.15,
		}, nil)
		value, err := FromBody.GetFloat(req, "weight")
		assert.Equal(t, 104.15, *value)
		assert.Nil(t, err)
	}
	{
		weight := float32(104.15)
		mockReq.EXPECT().BodyMap().Return(map[string]interface{}{
			"weight": weight,
		}, nil)
		value, err := FromBody.GetFloat(req, "weight")
		assert.Equal(t, float64(weight), *value)
		assert.Nil(t, err)
	}
}

func TestFromBody_GetBool(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockReq := NewMockRequest(ctrl)
	var req Request = mockReq
	{
		mockReq.EXPECT().BodyMap().Return(nil, fmt.Errorf("unknown error"))
		value, err := FromBody.GetBool(req, "agree")
		assert.Nil(t, value)
		assert.EqualError(t, err, "unknown error")
	}
	{
		mockReq.EXPECT().BodyMap().Return(nil, nil)
		value, err := FromBody.GetBool(req, "agree")
		assert.Nil(t, value)
		assert.Nil(t, err)
	}
	{
		mockReq.EXPECT().BodyMap().Return(map[string]interface{}{
			"agree": "abcd",
		}, nil)
		value, err := FromBody.GetBool(req, "agree")
		assert.Nil(t, value)
		assert.EqualError(t, err, "invalid 'agree', must be true or false")
	}
	{
		mockReq.EXPECT().BodyMap().Return(map[string]interface{}{
			"agree": "3465",
		}, nil)
		value, err := FromBody.GetBool(req, "agree")
		assert.Nil(t, value)
		assert.EqualError(t, err, "invalid 'agree', must be true or false")
	}
	{
		mockReq.EXPECT().BodyMap().Return(map[string]interface{}{
			"agree": 1231,
		}, nil)
		value, err := FromBody.GetBool(req, "agree")
		assert.Nil(t, value)
		assert.EqualError(t, err, "invalid 'agree', must be true or false")
	}
	{
		mockReq.EXPECT().BodyMap().Return(map[string]interface{}{
			"agree": true,
		}, nil)
		value, err := FromBody.GetBool(req, "agree")
		assert.Equal(t, true, *value)
		assert.Nil(t, err)
	}
	{
		mockReq.EXPECT().BodyMap().Return(map[string]interface{}{
			"agree": false,
		}, nil)
		value, err := FromBody.GetBool(req, "agree")
		assert.Equal(t, false, *value)
		assert.Nil(t, err)
	}
}
