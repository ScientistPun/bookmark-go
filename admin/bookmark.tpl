{{define "content"}}
<div class="page-title">
    <div class="title-env">
        <h1 class="title">导航管理</h1>
        <p class="description">导航与标签管理</p>
    </div>
</div>
<div class="row">
    <div class="col-md-12">
        <ul class="nav nav-pills bg-gray">
            <li class="active">
                <a href="#links" data-toggle="tab">标签&链接</a>
            </li>
            <li>
                <a href="#custom" data-toggle="tab">自定义</a>
            </li>
        </ul>
        <div class="tab-content">
            <div class="tab-pane active" id="links">
                <div class="panel">
                    <div class="panel-heading">添加标签</div>
                    <div class="panel-body">
                        <form class="row custom-form" id="tag-form">
                            <div class="col-xs-12 col-sm-6 col-md-4">
                                <div class="form-group row">
                                    <label class="col-sm-4 col-form-label">Icon</label>
                                    <div class="col-sm-8">
                                        <input type="text" class="form-control" name="icon" value=""/>
                                    </div>
                                </div>
                                <div class="form-group row">
                                    <label class="col-sm-4 col-form-label">标题</label>
                                    <div class="col-sm-8">
                                        <input type="text" class="form-control" name="title" value=""/>
                                    </div>
                                </div>
                                <div class="form-group row">
                                    <label class="col-sm-4 col-form-label">名称</label>
                                    <div class="col-sm-8">
                                        <input type="text" class="form-control" name="name" value=""/>
                                    </div>
                                </div>
                                <div class="form-group row">
                                    <label class="col-sm-4">&nbsp</label>
                                    <div class="col-sm-8"><input type="submit" class="btn btn-turquoise" value="添加"></div>
                                </div>
                            </div>
                        </form>
                    </div>
                </div>
                <div class="panel">
                    <div class="panel-heading">添加链接</div>
                    <div class="panel-body">
                        <form class="row custom-form" id="link-form">
                            <div class="col-xs-12 col-sm-6 col-md-4">
                                <div class="form-group row">
                                    <label class="col-sm-4 col-form-label">标题</label>
                                    <div class="col-sm-8">
                                        <input type="text" class="form-control" name="title" value=""/>
                                    </div>
                                </div>
                                <div class="form-group row">
                                    <label class="col-sm-4 col-form-label">路径</label>
                                    <div class="col-sm-8">
                                        <textarea class="form-control" name="url"></textarea>
                                    </div>
                                </div>
                                <div class="form-group row">
                                    <label class="col-sm-4 col-form-label">标签</label>
                                    <div class="col-sm-8">
                                        <select name="tag" class="form-control">
                                            {{range $k, $tag := .bookmark}}
                                            <option value="{{ $tag.Name }}">[{{ $tag.Name }}] {{ $tag.Title }}</option>
                                            {{end}}
                                        </select>
                                    </div>
                                </div>
                                <div class="form-group row">
                                    <label class="col-sm-4">&nbsp</label>
                                    <div class="col-sm-8"><input type="submit" class="btn btn-turquoise" value="添加"></div>
                                </div>
                            </div>
                        </form>
                    </div>
                </div>

                <div class="panel">
                    <div class="panel-heading">链接列表</div>
                    <div class="panel-body">
                        <table class="table table-hover">
                            <thead class="thead-light">
                            <tr>
                                <th scope="col">&nbsp;</th>
                                <th scope="col">标题</th>
                                <th scope="col">路径</th>
                                <th scope="col">说明</th>
                                <th scope="col">&nbsp;</th>
                            </tr>
                            </thead>
                            <tbody>
                            {{range $k, $tag := .bookmark}}
                                {{range $i, $link := $tag.Links}}
                                <tr>
                                    <td>[{{ $tag.Name }}] {{ $tag.Title }}</td>
                                    <td>
                                        <img src="{{ $link.Icon }}" width="32px">
                                        {{ $link.Title }}
                                    </td>
                                    <td>{{ $link.Url }} <span data-link="{{ $link.Url }}" class="btn btn-xs btn-warning copy-btn">复制</span></td>
                                    <td>{{ $link.Desc }}</td>
                                    <td>
                                        <span class="btn btn-sm btn-info refresh-link-btn" data-tag="{{ $tag.Name }}" data-sort="{{ $i }}"><i class="fa-refresh"></i> 更新</span>
                                        <a href="{{ $link.Url }}" class="btn btn-sm btn-secondary" target="_blank"><i class="fa-arrow-right"></i> 访问</a>
                                    </td>
                                </tr>
                                {{end}}
                            {{end}}
                            </tbody>
                        </table>
                    </div>
                </div>
            </div>
            <div class="tab-pane" id="custom">
                <div class="panel">
                    <div class="panel-heading">编辑yaml文件</div>
                    <div class="panel-body">
                        <form class="row custom-form" id="yaml-form">
                            <div class="col-xs-12 col-sm-10 col-md-8">
                                <div class="form-group row">
                                    <textarea class="form-control" name="content" cols="30" rows="35">{{ .bookmarkText }}</textarea>
                                </div>
                                <div class="form-group row">
                                    <label class="col-sm-4">&nbsp</label>
                                    <div class="col-sm-8">
                                        <input type="submit" class="btn btn-primary" value="保存">
                                        <input type="reset" class="btn" value="还原">
                                    </div>
                                </div>
                            </div>
                        </form>
                    </div>
                </div>
            </div>
        </div>
    </div>
</div>
<script>
    $(function () {
        $('#yaml-form').submit((e) => {
            let form = new FormData(document.getElementById("yaml-form"))
            $.ajax({
                url: "/manage/bookmark/edit-file",
                type: "post",
                data: form,
                processData: false,
                contentType: false,
                success: succFunc,
                error: errFunc
            })
            return false
        })
        $('#link-form').submit((e) => {
            let form = new FormData(document.getElementById("link-form"))
            $.ajax({
                url: "/manage/bookmark/add-link",
                type: "put",
                data: form,
                processData: false,
                contentType: false,
                success: succFunc,
                error: errFunc
            })
            return false
        })
        $('#tag-form').submit((e) => {
            let form = new FormData(document.getElementById("tag-form"))
            $.ajax({
                url: "/manage/bookmark/add-tag",
                type: "put",
                data: form,
                processData: false,
                contentType: false,
                success: succFunc,
                error: errFunc
            })
            return false
        })

        $(document).on('click', '.refresh-link-btn', function () {
            $.ajax({
                url: "/manage/bookmark/refresh-link/"+$(this).data('tag')+"/"+$(this).data('sort'),
                type: "get",
                processData: false,
                contentType: false,
                success: succFunc,
                error: errFunc
            })
            return false
        })

        $(document).on('click', '.copy-btn', function () {
            copyText($(this).data('link')) ? alert('复制成功!'):alert('操作错误!')
        })
    })
</script>
{{end}}