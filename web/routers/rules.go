/*

Copyright (c) 2018 sec.xiaomi.com

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THEq
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.

*/

package routers

import (
	"../../models"
	"../../util/common"
	"../../vars"
	"gopkg.in/macaron.v1"

	"github.com/go-macaron/session"
	"github.com/go-macaron/csrf"

	"strings"
	"strconv"
)

func ListRules(ctx *macaron.Context, sess session.Store) {
	page := ctx.Params(":page")
	p, _ := strconv.Atoi(page)
	p, pre, next := common.GetPreAndNext(p)

	if sess.Get("admin") != nil {
		rules, pages, _ := models.GetRulesPage(p)
		pageList := common.GetPageList(p, vars.PageStep, pages)

		ctx.Data["pages"] = pages
		ctx.Data["page"] = p
		ctx.Data["pre"] = pre
		ctx.Data["next"] = next
		ctx.Data["pageList"] = pageList
		ctx.Data["rules"] = rules
		ctx.HTML(200, "rules")
	} else {
		ctx.Redirect("/admin/login/")
	}
}

func NewRules(ctx *macaron.Context, sess session.Store) {
	if sess.Get("admin") != nil {
		ctx.HTML(200, "rules_new")
	} else {
		ctx.Redirect("/admin/login/")
	}
}

func DoNewRules(ctx *macaron.Context, sess session.Store) {
	ctx.Req.ParseForm()
	if sess.Get("admin") != nil {
		part := strings.TrimSpace(ctx.Req.Form.Get("part"))
		Type := strings.TrimSpace(ctx.Req.Form.Get("type"))
		content := strings.TrimSpace(ctx.Req.Form.Get("content"))
		caption := strings.TrimSpace(ctx.Req.Form.Get("caption"))
		desc := strings.TrimSpace(ctx.Req.Form.Get("desc"))
		status := strings.TrimSpace(ctx.Req.Form.Get("status"))
		intStatus, _ := strconv.Atoi(status)
		rule := models.NewRules(part, Type, content, caption, desc, intStatus)
		rule.Insert()
		ctx.Redirect("/admin/rules/list/")
	} else {
		ctx.Redirect("/admin/login/")
	}
}

func EditRules(ctx *macaron.Context, sess session.Store, x csrf.CSRF) {
	if sess.Get("admin") != nil {
		id := ctx.Params(":id")
		Id, _ := strconv.Atoi(id)
		rules, _, _ := models.GetRuleById(int64(Id))
		ctx.Data["csrf_token"] = x.GetToken()
		ctx.Data["rules"] = rules
		ctx.Data["user"] = sess.Get("admin")
		ctx.HTML(200, "rules_edit")
	} else {
		ctx.Redirect("/admin/login/")
	}
}

func DoEditRules(ctx *macaron.Context, sess session.Store) {
	ctx.Req.ParseForm()
	if sess.Get("admin") != nil {
		id := ctx.Params(":id")
		Id, _ := strconv.Atoi(id)
		part := strings.TrimSpace(ctx.Req.Form.Get("part"))
		Type := strings.TrimSpace(ctx.Req.Form.Get("type"))
		content := strings.TrimSpace(ctx.Req.Form.Get("content"))
		caption := strings.TrimSpace(ctx.Req.Form.Get("caption"))
		desc := strings.TrimSpace(ctx.Req.Form.Get("desc"))
		status := strings.TrimSpace(ctx.Req.Form.Get("status"))
		intStatus, _ := strconv.Atoi(status)
		models.EditRuleById(int64(Id), part, Type, content, caption, desc, intStatus)
		ctx.Redirect("/admin/rules/list/")
	} else {
		ctx.Redirect("/admin/login/")
	}
}

func DeleteRules(ctx *macaron.Context, sess session.Store) {
	if sess.Get("admin") != nil {
		id := ctx.Params(":id")
		Id, _ := strconv.Atoi(id)
		models.DeleteRulesById(int64(Id))
		ctx.Redirect("/admin/rules/list/")
	} else {
		ctx.Redirect("/admin/login/")
	}
}

func EnableRules(ctx *macaron.Context, sess session.Store) {
	if sess.Get("admin") != nil {
		id := ctx.Params(":id")
		Id, _ := strconv.Atoi(id)
		models.EnableRulesById(int64(Id))
		ctx.Redirect("/admin/rules/list/")
	} else {
		ctx.Redirect("/admin/login/")
	}
}

func DisableRules(ctx *macaron.Context, sess session.Store) {
	if sess.Get("admin") != nil {
		id := ctx.Params(":id")
		Id, _ := strconv.Atoi(id)
		models.DisableRulesById(int64(Id))
		ctx.Redirect("/admin/rules/list/")
	} else {
		ctx.Redirect("/admin/login/")
	}
}
