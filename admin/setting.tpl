{{define "content"}}
<div class="page-title">
    <div class="row">
        <div class="col-md-12">
            <div class="title-env">
                <h1 class="title">系统设置</h1>
                <p class="description">基础及用户设置</p>
            </div>
        </div>
    </div>
</div>
<div class="panel">
    <div class="panel-heading">基础设置</div>
    <div class="panel-body">
        <form class="row custom-form" id="system-form">
            <div class="col-xs-12 col-sm-6 col-md-4">
                <div class="form-group row">
                    <label for="port" class="col-sm-4 col-form-label">端口</label>
                    <div class="col-sm-8">
                        <input type="text" class="form-control" id="port" name="port" aria-describedby="portHelp" value="{{ .glob_conf.Service.Port }}"/>
                        <small id="portHelp" class="form-text text-muted">默认8080</small>
                    </div>
                </div>
                <div class="form-group row">
                    <label class="col-sm-4 col-form-label">皮肤</label>
                    <div class="col-sm-8">
                        <select name="skin" class="form-control">
                            {{range $k, $skin := .skins}}
                            <option value="{{ $skin }}" {{ if (eq $.useSkin $skin) }}selected=""{{end}}>{{ $skin }}</option>
                            {{end}}
                        </select>
                    </div>
                </div>
                <div class="form-group row">
                    <label class="col-sm-4 col-form-label">SSL访问</label>
                    <div class="col-sm-8">
                        <label for="ssl" class="checkbox-inline">
                            <input type="checkbox" id="ssl" name="ssl" aria-describedby="sslHelp" {{ if .glob_conf.Service.SSLMode }}checked=""{{end}} value="1"> 开启
                        </label>
                        <br/>
                        <small id="sslHelp" class="form-text text-muted">请确认ssl证书确实存在</small>
                    </div>
                </div>
                <div class="form-group row">
                    <label for="csr" class="col-sm-4 col-form-label">证书签名请求</label>
                    <div class="col-sm-8">
                        <input type="text" class="form-control" id="csr" name="csr" aria-describedby="csrHelp" value="{{ .glob_conf.Service.Csr }}"/>
                        <small id="csrHelp" class="form-text text-muted">ssl的csr地址</small>
                    </div>
                </div>
                <div class="form-group row">
                    <label for="pri-key" class="col-sm-4 col-form-label">私钥</label>
                    <div class="col-sm-8">
                        <input type="text" class="form-control" id="pri-key" name="key" aria-describedby="keyHelp" value="{{ .glob_conf.Service.Key }}"/>
                        <small id="keyHelp" class="form-text text-muted">ssl的私钥地址</small>
                    </div>
                </div>
                <div class="form-group row">
                    <label class="col-sm-4">&nbsp;</label>
                    <div class="col-sm-8"><input type="submit" class="btn btn-primary" value="保存"><p><i class="fa-info-circle"></i> 部分设置重启服务后生效</p></div>
                </div>
            </div>
        </form>
    </div>
</div>
<div class="panel">
    <div class="panel-heading">用户设置</div>
    <div class="panel-body">
        <form class="row custom-form" id="auth-form">
            <div class="col-xs-12 col-sm-6 col-md-4">
                <div class="form-group row">
                    <label for="password" class="col-sm-4 col-form-label">用户名</label>
                    <div class="col-sm-8">
                        <input type="text" class="form-control" name="username" value="{{ .glob_conf.Auth.Username }}">
                    </div>
                </div>
                <div class="form-group row">
                    <label for="password" class="col-sm-4 col-form-label">口令</label>
                    <div class="col-sm-8">
                        <input type="text" class="form-control" name="password" aria-describedby="pwdHelp" id="password">
                        <small id="pwdHelp" class="form-text text-muted">进入系统设置的口令，留空则不修改</small>
                    </div>
                </div>
                <div class="form-group row">
                    <label class="col-sm-4">&nbsp;</label>
                    <div class="col-sm-8"><button type="submit" class="btn btn-primary">保存</button></div>
                </div>
            </div>
        </form>
    </div>
</div>
<script>
    $(function () {
        $('#system-form').submit((e) => {
            var form = new FormData(document.getElementById("system-form"))
            $.ajax({
                url: "/manage/setting/basic",
                type: "post",
                data: form,
                processData:false,
                contentType:false,
                success: succFunc,
                error: errFunc
            })
            return false
        })

        $('#auth-form').submit((e) => {
            var form = new FormData(document.getElementById("auth-form"))
            $.ajax({
                url: "/manage/setting/auth",
                type: "post",
                data: form,
                processData:false,
                contentType:false,
                success: succFunc,
                error: errFunc
            })
            return false
        })
    })
</script>
{{end}}