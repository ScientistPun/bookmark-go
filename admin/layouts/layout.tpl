<!DOCTYPE html>
<html lang="ch">
<head>
    <meta charset="utf-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width, initial-scale=1.0"/>
    <title>{{ .title }}</title>
    <link rel="stylesheet" href="//fonts.googleapis.com/css?family=Arimo:400,700,400italic">
    <link rel="stylesheet" href="/manage/static/css/fonts/linecons/css/linecons.css">
    <link rel="stylesheet" href="/manage/static/css/fonts/fontawesome/css/font-awesome.min.css">
    <link rel="stylesheet" href="/manage/static/css/bootstrap.css">
    <link rel="stylesheet" href="/manage/static/css/xenon-core.css">
    <link rel="stylesheet" href="/manage/static/css/xenon-forms.css">
    <link rel="stylesheet" href="/manage/static/css/xenon-components.css">
    <link rel="stylesheet" href="/manage/static/css/xenon-skins.css">
    <link rel="stylesheet" href="/manage/static/css/backend.css">
    <link rel="shortcut icon" href="/manage/static/images/favicon.png" type="image/x-icon">
    <script src="/manage/static/js/jquery-1.11.1.min.js"></script>
    <!-- Imported styles on this page -->
    <!--    <link rel="stylesheet" href="/manage/static/js/dropzone/css/dropzone.css">-->
</head>
<body class="page-body">
    <div class="page-container">

        {{include "layouts/sidebar"}}

        <div class="main-content">
            {{include "layouts/navbar"}}

            {{template "content" .}}

            {{include "layouts/footer"}}
        </div>
    </div>
    <!-- Imported scripts on this page -->
    <!--<script src="/manage/static/js/dropzone/dropzone.min.js"></script>-->
    <!-- Bottom Scripts -->
    <script src="/manage/static/js/bootstrap.min.js"></script>
    <script src="/manage/static/js/TweenMax.min.js"></script>
    <script src="/manage/static/js/resizeable.js"></script>
    <script src="/manage/static/js/joinable.js"></script>
    <script src="/manage/static/js/xenon-api.js"></script>
    <script src="/manage/static/js/xenon-toggles.js"></script>
    <!-- JavaScripts initializations and stuff -->
    <script src="/manage/static/js/backend.js"></script>
</body>
<script>
    const succFunc = (ret) => {
        alert(ret.msg)
        location.reload()
    };
    const errFunc = (err) => {
        alert(err.responseJSON.msg)
        return false
    };
    const copyText = (text) => {
        var textarea = document.createElement("input");//创建input对象
        var currentFocus = document.activeElement;//当前获得焦点的元素
        document.body.appendChild(textarea);//添加元素
        textarea.value = text;
        textarea.focus();
        if(textarea.setSelectionRange)
            textarea.setSelectionRange(0, textarea.value.length); //获取光标起始位置到结束位置
        else
            textarea.select();
        try {
            var flag = document.execCommand("copy"); //执行复制
        } catch(eo) {
            var flag = false;
        }
        document.body.removeChild(textarea); //删除元素
        currentFocus.focus();
        return flag;
    }

    $(function () {
        $(document).on('#sign-out-btn', 'click', function () {
            $.get("/sign-out", function (ret) {
                if (ret.errCode != 0) {
                    alert(ret.errMsg)
                } else {
                    location.href = "/"
                }
                return false
            })
            alert("系统错误!!")
        })
    })
</script>
</html>