{{define "layouts/navbar"}}
<nav class="navbar user-info-navbar" role="navigation">
    <!-- User Info, Notifications and Menu Bar -->
    <!-- Left links for user info navbar -->
    <ul class="user-info-menu left-links list-inline list-unstyled">
        <li class="hidden-sm hidden-xs">
            <a href="#" data-toggle="sidebar">
                <i class="fa-bars"></i>
            </a>
        </li>
        <li class="hidden-sm hidden-xs">
            <a href="/"> 主页 </a>
        </li>
    </ul>
    <!-- Right links for user info navbar -->
    <ul class="user-info-menu right-links list-inline list-unstyled">
        <li class="dropdown user-profile">
            <a href="/sign-out" id="sign-out-btn">
                <i class="fa-lock"></i> 退出
            </a>
        </li>
    </ul>
</nav>
{{end}}