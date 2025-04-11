package restaurant_service

import (
	restaurant_model "Food-Delivery/internal/restaurant/entity/dto"
	"context"
	"testing"
)

type mockRepository interface {
	Create(ctx context.Context, dto *restaurant_model.RestaurantCreateDTO) error
}

func TestCreateRestaurant(t *testing.T) {
	//repo := restaurant_repository.NewRestaurantRepository(db)
	//
	//service := NewRestaurantService(repo)
	//
	////Test những trường hợp lỗi
	//dataTable := []struct {
	//	input    entity.Restaurant
	//	expected string
	//}{
	//	{input: entity.Restaurant{Name: "", Address: "Nguyễn Huệ, Q1"}, expected: entity.ErrNameIsEmpty.Error()},
	//	{input: entity.Restaurant{Name: "Candy Home", Address: ""}, expected: entity.ErrAddressIsEmpty.Error()},
	//}
	//
	//for _, item := range dataTable {
	//	err := service.Create(context.Background(), &item.input)
	//	fmt.Println(err)
	//	if err.Error() != item.expected {
	//		t.Errorf("create place - Input %v, Expected: %v, Output: %v", item.input, item.expected, err)
	//	}
	//}
	//
	////Test trường hợp thành công, expect là không trả về lỗi
	//dataTest := entity.Restaurant{Name: "SweetHome", Address: "CMT8, Q3"}
	//err := service.Create(context.Background(), &dataTest)
	//if err != nil {
	//	t.Errorf("create place - Input: %v, Output: %v", dataTest, err)
	//}

}
