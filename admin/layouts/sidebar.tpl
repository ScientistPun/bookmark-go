{{define "layouts/sidebar"}}
<div class="sidebar-menu toggle-others fixed">
    <div class="sidebar-menu-inner">
        <header class="logo-env">
            <!-- logo -->
            <div class="logo">
                <a href="index.html" class="logo-expanded">
                    <img src="./static/images/logo@2x.png" width="100%" alt="" />
                </a>
                <a href="index.html" class="logo-collapsed">
                    <img src="./static/images/logo-collapsed@2x.png" width="40" alt="" />
                </a>
            </div>
        </header>
        <ul id="main-menu" class="main-menu">
            <!-- add class "multiple-expanded" to allow multiple submenus to open -->
            <!-- class "auto-inherit-active-class" will automatically add "active" class for parent elements who are marked already with class "active" -->
            {{range .glob_nav}}
            <li {{if (eq .Url $.uri) }}class="active"{{end}}>
                <a href="{{ .Url }}">
                    <i class="{{ .Icon }}"></i>
                    <span class="title">{{ .Title }}</span>
                </a>
            </li>
            {{end}}
<!--            <li>-->
<!--                <a href="tags">-->
<!--                    <i class="linecons-database"></i>-->
<!--                    <span class="title">导航管理</span>-->
<!--                </a>-->
<!--            </li>-->
        </ul>
    </div>
</div>
{{end}}