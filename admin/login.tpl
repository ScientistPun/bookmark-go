<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>{{ .title }}</title>

</head>
<body>
</body>
<script src="/manage/static/js/jquery-1.11.1.min.js"></script>
<script src="/manage/static/lib/layer/layer.js"></script>
<script>
    layer.open({
        type: 1,
        area: ['420px', '380px'],
        title: "请输入用户名和口令",
        skin: 'layui-layer-rim', //样式类名
        content: '<form id="sign-in-form" method="post">\n' +
            '    <table cellpadding="4" cellspacing="8">\n' +
            '        <tr>\n' +
            '            <td>用户名</td>\n' +
            '            <td><input type="text" name="username"></td>\n' +
            '        </tr>\n' +
            '        <tr>\n' +
            '            <td>口令</td>\n' +
            '            <td><input type="password" name="password"></td>\n' +
            '        </tr>\n' +
            '        <tr align="center">\n' +
            '            <td colspan="2"><input type="submit" value="提交"></td>\n' +
            '        </tr>\n' +
            '    </table>\n' +
            '</form>'
    });

    $(function () {
        $(document).on('submit', '#sign-in-form', function () {
            let form = new FormData(this)
            $.ajax({
                url: "/sign-in",
                type: "post",
                data: form,
                processData:false,
                contentType:false,
                success: (ret) => {
                    location.reload()
                },
                error: (err) => {
                    alert(err.responseJSON.msg)
                }
            })
            return false
        })
    })
</script>
</html>