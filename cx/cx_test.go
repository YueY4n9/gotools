package cx

import (
	"fmt"
	"github.com/YueY4n9/gotools/echo"
	"net/http"
	"net/url"
	"testing"
	"time"
)

func TestSendMsg(t *testing.T) {
	appKeys := []string{
		"4bd910aa48803803e4d7a2ed334ed436-127k063d", // peiran
		//"f043a5677976552b348492635eb2e1a0-ujJ3xRV2", // wufei
	}
	userInfoStr := "{\"user_id\":\"3291738c\",\"system_fields\":{\"name\":\"岳杨\",\"en_name\":\"\",\"email\":\"yueyang@manjuu.com\",\"mobile\":\"+86 17806283435\",\"department_id\":\"od-874befe21984374d1a312fdfe8310c3d\",\"manager\":{\"user_id\":\"9428c5g7\",\"name\":\"陶晟\",\"en_name\":\"陶晟\"},\"gender\":1,\"birthday\":\"1996-11-18\",\"native_region\":{\"iso_code\":\"CHN-37\",\"name\":\"山东省\"},\"ethnicity\":1,\"marital_status\":0,\"political_status\":12,\"entered_workforce_date\":\"2019-07-01\",\"id_type\":1,\"id_number\":\"371082199611180013\",\"hukou_type\":2,\"hukou_location\":\"山东省威海市\",\"bank_account_number\":\"6214831057159813\",\"bank_name\":\"\",\"social_security_account\":\"\",\"provident_fund_account\":\"\",\"employee_no\":\"\",\"employee_type\":1,\"status\":2,\"hire_date\":\"2023-10-10\",\"probation_months\":6,\"conversion_date\":\"2024-04-10\",\"application\":0,\"application_status\":3,\"last_day\":\"\",\"departure_type\":0,\"departure_reason\":0,\"departure_notes\":\"\",\"contract_company\":{\"id\":7324625332804108316,\"name\":\"蛮啾\"},\"contract_type\":1,\"contract_start_date\":\"2023-10-10\",\"contract_expiration_date\":\"2026-10-09\",\"contract_sign_times\":1,\"personal_email\":\"mmdyueyang@gmail.com\",\"family_address\":\"上海市浦东新区培元新苑5号楼803\",\"primary_emergency_contact\":{\"name\":\"岳丽强\",\"relationship\":1,\"mobile\":\"13906315939\"},\"emergency_contact\":[{\"name\":\"岳丽强\",\"relationship\":1,\"mobile\":\"13906315939\"}],\"highest_level_of_edu\":{\"level\":7,\"school\":\"山东科技大学\",\"major\":\"计算机科学与技术\",\"degree\":1,\"start\":\"2015-09-01\",\"end\":\"2019-06-30\"},\"education\":[{\"level\":7,\"school\":\"山东科技大学\",\"major\":\"计算机科学与技术\",\"degree\":1,\"start\":\"2015-09-01\",\"end\":\"2019-06-30\"}],\"id_photo_po_side\":[{\"id\":\"MGinbLhTHoNp9oxPsvacYFZVnif\",\"mime_type\":\"jpg\",\"name\":\"471695378069_.pic.jpg\",\"size\":267110}],\"id_photo_em_side\":[{\"id\":\"Ft7Cbv7FDonLoixrPtBcfWuWnuh\",\"mime_type\":\"jpg\",\"name\":\"461695378068_.pic.jpg\",\"size\":327301}],\"diploma_photo\":[{\"id\":\"NVZvb1CHQoo8oyxzuZlcLY6Xnsl\",\"mime_type\":\"jpg\",\"name\":\"521695378440_.pic.jpg\",\"size\":492651},{\"id\":\"JXlybxGDdonFA7xAarlcpFeon3f\",\"mime_type\":\"jpg\",\"name\":\"531695378441_.pic.jpg\",\"size\":222280}],\"graduation_cert\":[{\"id\":\"OiODb8eoPouHmvxQfIKc4DSAnaf\",\"mime_type\":\"jpg\",\"name\":\"501695378439_.pic.jpg\",\"size\":427926},{\"id\":\"GFJxbTYtwoC9RlxysbmcxlignOx\",\"mime_type\":\"jpg\",\"name\":\"511695378439_.pic.jpg\",\"size\":171540}],\"offboarding_file\":[{\"id\":\"ZkVHbD4NjociuyxMu6WckiYmnkh\",\"mime_type\":\"jpg\",\"name\":\"541695378558_.pic.jpg\",\"size\":191336}],\"cancel_onboarding_reason\":0,\"cancel_onboarding_notes\":\"\",\"employee_form_status\":3,\"create_time\":1695373066000,\"update_time\":1720065520000},\"custom_fields\":[{\"key\":\"custom_7325338999987553819\",\"label\":\"是否有海外工作经历\",\"type\":\"enum\",\"value\":\"{\\\"key\\\":\\\"E-7325343735235117084\\\",\\\"value\\\":\\\"否\\\"}\"},{\"key\":\"custom_7325338999987914267\",\"label\":\"海外工作经历累计年限（单位：年）\",\"type\":\"text\"},{\"key\":\"custom_7325340182256387610\",\"label\":\"海外工作经历国家\",\"type\":\"text\"},{\"key\":\"customField_1677813250154\",\"label\":\"公积金账号（上海市）\",\"type\":\"text\",\"value\":\"无\"},{\"key\":\"customField_1675409601291\",\"label\":\"补充公积金账号（上海）\",\"type\":\"text\"},{\"key\":\"customField_1677813289306\",\"label\":\"开户行支行信息\",\"type\":\"text\",\"value\":\"招商银行北京分行北京圆明园西路支行\"},{\"key\":\"customField_1674909885894\",\"label\":\"工资流水或社保明细\",\"type\":\"attachment\",\"value\":\"[{\\\"id\\\":\\\"JqXHbTBVNos5wjxzZSJc4k42nAd\\\",\\\"mime_type\\\":\\\"jpg\\\",\\\"name\\\":\\\"551695378590_.pic.jpg\\\",\\\"size\\\":189656}]\"},{\"key\":\"custom_7314625014498084415\",\"label\":\"个人形象照片（用于制作工牌）\",\"type\":\"attachment\",\"value\":\"[]\"},{\"key\":\"custom_7325732934459246118\",\"label\":\"职业等级证书\",\"type\":\"attachment\",\"value\":\"[]\"},{\"key\":\"custom_7325325770322527763\",\"label\":\"其他\",\"type\":\"attachment\",\"value\":\"[]\"}]}"
	title := "yueyang" + " is resigned"
	content := fmt.Sprintf("姓名: %v\n时间: %v\n描述: %v\n", "yueyang", time.Now().Format("2006-01-02 15:04:05"), userInfoStr)
	for _, appKey := range appKeys {
		time.Sleep(2 * time.Second)
		params := url.Values{}
		params.Add("appkey", appKey)
		params.Add("title", title)
		params.Add("content", content)
		resp, err := http.Get("https://cx.super4.cn/push_msg" + "?" + params.Encode())
		if err != nil {
			fmt.Println(err)
		}
		defer resp.Body.Close()
		echo.Json(resp.Status)
	}
}
