{{define "content"}}
<div class="page-title">
    <div class="title-env">
        <h1 class="title">搜索引擎</h1>
        <p class="description">搜索引擎文件设置</p>
    </div>
</div>
<div class="panel">
    <div class="panel-heading">编辑搜索文件</div>
    <div class="panel-body">
        <form class="row custom-form" id="yaml-form">
            <div class="col-xs-12 col-sm-10 col-md-8">
                <div class="form-group row">
                    <textarea class="form-control" name="content" cols="30" rows="35">{{ .searchEngineText }}</textarea>
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
<script>
    $(function () {
        $('#yaml-form').submit((e) => {
            let form = new FormData(document.getElementById("yaml-form"))
            $.ajax({
                url: "/manage/search/edit",
                type: "post",
                data: form,
                processData: false,
                contentType: false,
                success: succFunc,
                error: errFunc
            })
            return false
        })
    })
</script>
{{end}}