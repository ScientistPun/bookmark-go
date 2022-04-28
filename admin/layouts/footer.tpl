{{define "layouts/footer"}}
    <!-- Main Footer -->
    <!-- Choose between footer styles: "footer-type-1" or "footer-type-2" -->
    <!-- Add class "sticky" to  always stick the footer to the end of page (if page contents is small) -->
    <!-- Or class "fixed" to  always fix the footer to the end of page -->
    <footer class="main-footer fixed footer-type-1">
        <div class="footer-inner">
            <!-- Add your copyright text here -->
            <div class="footer-text">
                &copy; Bookmark Ver{{ .glob_conf.System.Version }} Design by <a href="http://webstack.cc"><strong>Webstack</strong></a>(<a href="http://viggoz.com" target="_blank"><strong>Viggo</strong></a>) Build by ScientistPun
            </div>
            <!-- Go to Top Link, just add rel="go-top" to any link to add this functionality -->
            <div class="go-up">
                <a href="#" rel="go-top">
                    <i class="fa-angle-up"></i>
                </a>
            </div>
        </div>
    </footer>
{{end}}