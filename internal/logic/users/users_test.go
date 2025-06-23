package users

import (
	"context"
	"fmt"
	_ "github.com/go-sql-driver/mysql"              // 必须匿名导入！  database/sql 驱动
	_ "github.com/gogf/gf/contrib/drivers/mysql/v2" // GoFrame ORM adapter
	"scaffold/internal/dao"
	"scaffold/internal/model"
	"scaffold/internal/service"
	"testing"
)

//todo 完成Mock实现
// ---- Mock 实现 ----
//type mockDAO struct {
//	scanErr      error
//	countReturn  int
//	insertErr    error
//	rowsAffected int64
//	lastInsertId int64
//}
//
//func (m *mockDAO) ScanAndCount(ctx context.Context, list interface{}, count *int, lock bool) error {
//	*count = m.countReturn
//	return m.scanErr
//}
//func (m *mockDAO) Insert(ctx context.Context, item interface{}) (Result, error) {
//	if m.insertErr != nil {
//		return nil, m.insertErr
//	}
//	return m, nil
//}
//func (m *mockDAO) RowsAffected() (int64, error) { return m.rowsAffected, nil }
//func (m *mockDAO) LastInsertId() (int64, error) { return m.lastInsertId, nil }
//
//type mockJwt struct {
//	saltErr error
//	saltVal string
//}
//
//func (m *mockJwt) GenerateSalt() (string, error)   { return m.saltVal, m.saltErr }
//func (m *mockJwt) HashPassword(p, s string) string { return fmt.Sprintf("H[%s|%s]", p, s) }
//
//// ---- 测试用例 ----
//func TestCreateUser_Success(t *testing.T) {
//	dao := &mockDAO{countReturn: 7, rowsAffected: 1, lastInsertId: 88}
//	jwt := &mockJwt{saltVal: "SALT"}
//	svc := NewUserService(dao, jwt)
//
//	in := model.UserSignUp{UserID: "u100", Username: "alice", Password: "pwd"}
//	reply, err := svc.CreateUser(context.Background(), in)
//	if err != nil {
//		t.Fatalf("expected no error, got %v", err)
//	}
//	want := "100000007"
//	if reply.UID != want {
//		t.Errorf("UID = %s; want %s", reply.UID, want)
//	}
//}
//
//func TestCreateUser_ScanError(t *testing.T) {
//	dao := &mockDAO{scanErr: fmt.Errorf("scan fail")}
//	svc := NewUserService(dao, &mockJwt{})
//	if _, err := svc.CreateUser(context.Background(), model.UserSignUp{}); err == nil ||
//		!strings.Contains(err.Error(), "scan fail") {
//		t.Errorf("expected scan error, got %v", err)
//	}
//}
//
//func TestCreateUser_GenerateSaltError(t *testing.T) {
//	dao := &mockDAO{countReturn: 0}
//	jwt := &mockJwt{saltErr: fmt.Errorf("salt fail")}
//	svc := NewUserService(dao, jwt)
//	if _, err := svc.CreateUser(context.Background(), model.UserSignUp{Password: "x"}); err == nil ||
//		!strings.Contains(err.Error(), "生成盐值失败") {
//		t.Errorf("expected salt error, got %v", err)
//	}
//}

// 你可以继续在这个文件里补充更多场景测试，比如：
// - 插入返回 err
// - RowsAffected() == 0
// - LastInsertId 返回 err
func TestCreateUser_Integration(t *testing.T) {
	service.Migrations().Migrate(context.Background())
	_, err := User().CreateUser(context.Background(), model.UserSignUp{
		UserID:   "admin",
		Username: "admin",
		Password: "152123123",
	})
	if err != nil {
		t.Fatalf("CreateUser failed: %v", err)
	}

	// 5. 查询数据库，断言结果
	record, err := dao.UserModels.Ctx(context.Background()).Count()
	if err != nil {
		t.Fatalf("fetch created user failed: %v", err)
	}
	if record != 1 {
		t.Errorf("count should == 1 , got %d", record)
	}
}

func TestValidateUser_Integration(t *testing.T) {
	service.Migrations().Migrate(context.Background())
	result, err := User().ValidateUser(context.Background(), model.UserSignIn{
		UserID:   "admin",
		Password: "admin5588",
	})
	fmt.Println(result)

	if err != nil {
		t.Fatalf("CreateUser failed: %v", err)
	}

	// 5. 查询数据库，断言结果
	record, err := dao.UserModels.Ctx(context.Background()).Count()
	if err != nil {
		t.Fatalf("fetch created user failed: %v", err)
	}
	if record != 1 {
		t.Errorf("count should == 1 , got %d", record)
	}
}
func TestValidateUser_Token(t *testing.T) {

	token, err := service.Jwt().GenerateToken(context.Background(), "admin")
	if err != nil {
		t.Fatalf("CreateUser failed: %v", err)
	}
	userID, err := service.Jwt().ParseToken(token)
	if err != nil {
		t.Fatalf("CreateUser failed: %v", err)
	}
	println(userID)
}
