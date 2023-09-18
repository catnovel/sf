package sf

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"github.com/catnovelapi/BuilderHttpClient"
	"github.com/tidwall/gjson"
)

func (s *Sfacg) GetBookInfoApi(bookId any) gjson.Result {
	return s.get(fmt.Sprintf("/novels/%v", bookId), map[string]any{"expand": "intro,sysTags"})
}

func (s *Sfacg) GetAccountInApi() gjson.Result {
	return s.get("/user", nil)
}

func (s *Sfacg) AccountInMoneyApi() gjson.Result {
	return s.post("/user/money", nil).Gjson()
}

func (s *Sfacg) LoginApi(username, password string) BuilderHttpClient.ResponseInterfaceBuilder {
	return s.post("/sessions", &loginPayload{Username: username, Password: password})
}
func (s *Sfacg) SearchNovelsResultApi(keyword string, page int) gjson.Result {
	return s.get("/search/novels/result", map[string]any{"q": keyword, "page": page, "size": "10"})
}

func (s *Sfacg) GetChapterInfoApi(chapterId any) string {
	for i := 0; i < 5; i++ {
		response := s.get(fmt.Sprintf("/Chaps/%v", chapterId), map[string]string{"expand": "content"})
		content := response.Get("data.expand.content").String()
		if content != "" {
			return content
		}
	}
	return ""
}
func (s *Sfacg) ChapterListByBookIDApi(bookId any) gjson.Result {
	return s.get(fmt.Sprintf("/novels/%s/dirs", bookId), map[string]string{"expand": "originNeedFireMoney"})
}
func (s *Sfacg) NewChapterListByBookIDApi(bookId any) []gjson.Result {
	var chapterListArray []gjson.Result
	for _, i := range s.ChapterListByBookIDApi(bookId).Get("data.volumeList").Array() {
		for _, j := range i.Get("chapterList").Array() {
			chapterListArray = append(chapterListArray, j)
		}
	}
	return chapterListArray
}

func (s *Sfacg) GetBookShelfApi() gjson.Result {
	return s.get("/user/Pockets", map[string]string{"expand": "novels"})
}

func (s *Sfacg) BookListApi(bookId any) gjson.Result {
	return s.get(fmt.Sprintf("/novels/%v/bookList", bookId), map[string]string{"size": "3", "page": "0"})
}

func (s *Sfacg) UpdateBooksList(page int) gjson.Result {
	return s.get("/novels", map[string]any{"page": page, "size": "50", "filter": "latest-signnvip", "expand": "sysTags,intro"})
}

func (s *Sfacg) AdpworksApi(bookId any) gjson.Result {
	return s.get(fmt.Sprintf("/adpworks/novelId/%v", bookId), map[string]string{"expand": "signlevel"})
}

func (s *Sfacg) GetPositionApi() gjson.Result {
	return s.get("/position", nil)
}

func (s *Sfacg) GetSpecialPushApi() gjson.Result {
	return s.get("/specialpush", map[string]string{"pushNames": "merchPush", "entityId": "", "entityType": "novel"})
}

func (s *Sfacg) GetWelfareCfgApi() gjson.Result {
	return s.get("/welfare/cfg", nil)
}

func (s *Sfacg) GetStaticsResourceApi() gjson.Result {
	return s.get("/StaticsResource", nil)
}

func (s *Sfacg) GetUserWelfareStoreitemsLatestApi() gjson.Result {
	return s.get("/user/welfare/storeitems/latest", nil)
}

func (s *Sfacg) essaySolicitationNovelApi(tagIds, page int) gjson.Result {
	return s.getWeb("https://pages.sfacg.com/api/essay/getNovels", map[string]any{"tagIds": tagIds, "page": page, "size": "50"}, true)
}

func (s *Sfacg) EssayShortNovelApi(page int) gjson.Result {
	return s.essaySolicitationNovelApi(655, page)
}

func (s *Sfacg) EssayNovellaApi(page int) gjson.Result {
	return s.essaySolicitationNovelApi(654, page)
}

func (s *Sfacg) EssayLongNovelApi(page int) gjson.Result {
	return s.essaySolicitationNovelApi(653, page)
}

func (s *Sfacg) SystemRecommendApi() gjson.Result {
	return s.get("/novel/systemRecommend", nil)
}

func (s *Sfacg) GetActConfigApi() gjson.Result {
	return s.getWeb("https://pages.sfacg.com/api/apptabproject/getActConfig", nil, false)
}
func (s *Sfacg) PostConversionsApi() BuilderHttpClient.ResponseInterfaceBuilder {
	randomBytes := make([]byte, 16)
	_, _ = rand.Read(randomBytes)
	return s.post("/androiddeviceinfos/conversion", putReadingTimePayload{OaID: hex.EncodeToString(randomBytes)})
}
func (s *Sfacg) VersionInformation() BuilderHttpClient.ResponseInterfaceBuilder {
	return s.post("/androidcfg", nil)
}
func (s *Sfacg) PreOrderApi() BuilderHttpClient.ResponseInterfaceBuilder {

	return s.post("/preOrder", map[string]any{"expand": "intro,typeName,tags,sysTags", "withExpiredPreOrder": false})
}
func (s *Sfacg) PostSpecialPushApi() BuilderHttpClient.ResponseInterfaceBuilder {
	return s.post("/specialpush", putSignInfoPayload{SignDate: "merchPush", EntityId: "", EntityType: "novel"})
}
