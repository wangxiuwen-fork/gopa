// Generated by ego.
// DO NOT EDIT

package common

import (
	"fmt"
	"html"
	"io"
)

var _ = fmt.Sprint("") // just so that we can keep the fmt import for now
func Body(w io.Writer) error {
	_, _ = io.WriteString(w, "\n\n</head>\n<body class=\"tm-background\">\n")
	return nil
}
func Footer(w io.Writer) error {
	_, _ = io.WriteString(w, "\n\n<div class=\"tm-footer\">\n    <div class=\"uk-container uk-container-center uk-text-center\">\n        <br>\n        <hr class=\"uk-article-divider\" >\n        <ul class=\"uk-subnav uk-subnav-line uk-flex-center\">\n            <li><a href=\"http://github.com/medcl/gopa\">GitHub</a></li>\n            <li><a href=\"http://github.com/medcl/gopa/issues\">Issues</a></li>\n            <li><a href=\"https://github.com/medcl/gopa/releases\">Releases</a></li>\n            <li><a href=\"http://github.com/medcl/gopa/blob/master/CHANGES.md\">Changelog</a></li>\n        </ul>\n\n        <div class=\"uk-panel\">\n            <p>Licensed under <a target=\"_blank\" href=\"https://github.com/medcl/gopa/blob/master/LICENSE\">Apache License, Version 2.0</a>.</p>\n            <a href=\"/admin/\"><svg  height=\"12\" width=\"64\">\n                <image xlink:href=\"/static/assets/img/logo.svg\" src=\"/static/assets/img/logo.svg\" height=\"12\" />\n            </svg></a>\n        </div>\n\n    </div>\n</div>\n</body>\n</html>\n")
	return nil
}
func Head(w io.Writer, title string, customHeaderBlock string) error {
	_, _ = io.WriteString(w, "\n<!DOCTYPE html>\n<html>\n<head>\n  <title>")
	_, _ = io.WriteString(w, html.EscapeString(fmt.Sprint(title)))
	_, _ = io.WriteString(w, " - GOPA</title>\n  <link rel=\"stylesheet\" href=\"/static/assets/uikit-2.27.1/css/uikit.min.css\" />\n  <link rel=\"icon\" href=\"/static/assets/img/favicon.ico\" type=\"image/x-icon\" />\n  <link rel=\"shortcut icon\" href=\"/static/assets/img/favicon.ico\" type=\"image/x-icon\" />\n  <script src=\"/static/assets/js/jquery.min.js\"></script>\n  <script src=\"/static/assets/uikit-2.27.1/js/uikit.min.js\"></script>\n\n  <meta charset=\"utf-8\">\n  ")
	if len(customHeaderBlock) > 0 {
		_, _ = io.WriteString(w, "\n  ")
		_, _ = io.WriteString(w, html.EscapeString(fmt.Sprint(customHeaderBlock)))
		_, _ = io.WriteString(w, "\n  ")
	}
	_, _ = io.WriteString(w, "\n")
	return nil
}
func Nav(w io.Writer, current string) error {
	_, _ = io.WriteString(w, "\n\n<nav class=\"tm-navbar uk-navbar uk-navbar-attached\">\n  <div class=\"uk-container uk-container-center\">\n    ")

	logoUrl := "/admin/"
	if len(navs) > 0 {
		logoUrl = navs[0].url
	}

	_, _ = io.WriteString(w, "\n    <a class=\"uk-navbar-brand uk-hidden-small\" href=\"")
	_, _ = io.WriteString(w, html.EscapeString(fmt.Sprint(logoUrl)))
	_, _ = io.WriteString(w, "\"><img class=\"uk-margin uk-margin-remove\" src=\"/static/assets/img/logo.svg\" height=\"30\" title=\"GOPA\" alt=\"GOPA\"></a>\n\n    <ul class=\"uk-navbar-nav uk-hidden-small\">\n\n      ")
	if len(navs) > 0 {
		for _, obj := range navs {

			_, _ = io.WriteString(w, "\n      <li ")
			_, _ = fmt.Fprint(w, NavCurrent(current, obj.name))
			_, _ = io.WriteString(w, " ><a href=\"")
			_, _ = io.WriteString(w, html.EscapeString(fmt.Sprint(obj.url)))
			_, _ = io.WriteString(w, "\">")
			_, _ = io.WriteString(w, html.EscapeString(fmt.Sprint(obj.displayName)))
			_, _ = io.WriteString(w, "</a></li>\n      ")

		}
	}

	_, _ = io.WriteString(w, "\n    </ul>\n\n    <a href=\"#tm-offcanvas\" class=\"uk-navbar-toggle uk-visible-small\" data-uk-offcanvas=\"\"></a>\n\n    <div class=\"uk-navbar-brand uk-navbar-center uk-visible-small\"><img src=\"/static/assets/img/logo.svg\" height=\"30\" title=\"GOPA\" alt=\"GOPA\"></div>\n\n  </div>\n</nav>\n\n<div style=\"height: 15px;clear: both\"></div>\n")
	return nil
}
func OffCanvas(w io.Writer) error {
	_, _ = io.WriteString(w, "\n<div id=\"tm-offcanvas\" class=\"uk-offcanvas\">\n\n  <div class=\"uk-offcanvas-bar\">\n\n    <ul class=\"uk-nav uk-nav-offcanvas uk-nav-parent-icon\" data-uk-nav=\"{multiple:true}\">\n      <li class=\"uk-parent uk-active uk-open\" aria-expanded=\"true\"><a href=\"#\">Documentation</a>\n        <div style=\"overflow: hidden; height: auto; position: relative;\"><ul class=\"uk-nav-sub\" role=\"menu\">\n          <li><a href=\"documentation_get-started.html\">Get started</a></li>\n          <li><a href=\"documentation_how-to-customize.html\">How to customize</a></li>\n          <li><a href=\"documentation_layouts.html\">Layout examples</a></li>\n          <li><a href=\"core.html\">Core</a></li>\n          <li><a href=\"components.html\">Components</a></li>\n          <li><a href=\"documentation_project-structure.html\">Project structure</a></li>\n          <li><a href=\"documentation_less-sass.html\">Less &amp; Sass files</a></li>\n          <li><a href=\"documentation_create-a-theme.html\">Create a theme</a></li>\n          <li><a href=\"documentation_create-a-style.html\">Create a style</a></li>\n          <li><a href=\"documentation_customizer-json.html\">Customizer.json</a></li>\n          <li><a href=\"documentation_javascript.html\">JavaScript</a></li>\n          <li><a href=\"documentation_custom-prefix.html\">Custom prefix</a></li>\n        </ul></div>\n      </li>\n      <li class=\"uk-nav-header\">Core</li>\n      <li class=\"uk-parent\" aria-expanded=\"false\"><a href=\"#\"><i class=\"uk-icon-wrench\"></i> Defaults</a>\n        <div style=\"overflow:hidden;height:0;position:relative;\" class=\"uk-hidden\"><ul class=\"uk-nav-sub\" role=\"menu\">\n          <li><a href=\"base.html\">Base</a></li>\n          <li><a href=\"print.html\">Print</a></li>\n        </ul></div>\n      </li>\n      <li class=\"uk-parent\" aria-expanded=\"false\"><a href=\"#\"><i class=\"uk-icon-th-large\"></i> Layout</a>\n        <div style=\"overflow:hidden;height:0;position:relative;\" class=\"uk-hidden\"><ul class=\"uk-nav-sub\" role=\"menu\">\n          <li><a href=\"grid.html\">Grid</a></li>\n          <li><a href=\"panel.html\">Panel</a></li>\n          <li><a href=\"block.html\">Block</a></li>\n          <li><a href=\"article.html\">Article</a></li>\n          <li><a href=\"comment.html\">Comment</a></li>\n          <li><a href=\"utility.html\">Utility</a></li>\n          <li><a href=\"flex.html\">Flex</a></li>\n          <li><a href=\"cover.html\">Cover</a></li>\n        </ul></div>\n      </li>\n      <li class=\"uk-parent\" aria-expanded=\"false\"><a href=\"#\"><i class=\"uk-icon-bars\"></i> Navigations</a>\n        <div style=\"overflow:hidden;height:0;position:relative;\" class=\"uk-hidden\"><ul class=\"uk-nav-sub\" role=\"menu\">\n          <li><a href=\"nav.html\">Nav</a></li>\n          <li><a href=\"navbar.html\">Navbar</a></li>\n          <li><a href=\"subnav.html\">Subnav</a></li>\n          <li><a href=\"breadcrumb.html\">Breadcrumb</a></li>\n          <li><a href=\"pagination.html\">Pagination</a></li>\n          <li><a href=\"tab.html\">Tab</a></li>\n          <li><a href=\"thumbnav.html\">Thumbnav</a></li>\n        </ul></div>\n      </li>\n      <li class=\"uk-parent\" aria-expanded=\"false\"><a href=\"#\"><i class=\"uk-icon-check\"></i> Elements</a>\n        <div style=\"overflow:hidden;height:0;position:relative;\" class=\"uk-hidden\"><ul class=\"uk-nav-sub\" role=\"menu\">\n          <li><a href=\"list.html\">List</a></li>\n          <li><a href=\"description-list.html\">Description list</a></li>\n          <li><a href=\"table.html\">Table</a></li>\n          <li><a href=\"form.html\">Form</a></li>\n        </ul></div>\n      </li>\n      <li class=\"uk-parent\" aria-expanded=\"false\"><a href=\"#\"><i class=\"uk-icon-folder-open\"></i> Common</a>\n        <div style=\"overflow:hidden;height:0;position:relative;\" class=\"uk-hidden\"><ul class=\"uk-nav-sub\" role=\"menu\">\n          <li><a href=\"button.html\">Button</a></li>\n          <li><a href=\"icon.html\">Icon</a></li>\n          <li><a href=\"close.html\">Close</a></li>\n          <li><a href=\"badge.html\">Badge</a></li>\n          <li><a href=\"alert.html\">Alert</a></li>\n          <li><a href=\"thumbnail.html\">Thumbnail</a></li>\n          <li><a href=\"overlay.html\">Overlay</a></li>\n          <li><a href=\"text.html\">Text</a></li>\n          <li><a href=\"column.html\">Column</a></li>\n          <li><a href=\"animation.html\">Animation</a></li>\n          <li><a href=\"contrast.html\">Contrast</a></li>\n        </ul></div>\n      </li>\n      <li class=\"uk-parent\" aria-expanded=\"false\"><a href=\"#\"><i class=\"uk-icon-magic\"></i> JavaScript</a>\n        <div style=\"overflow:hidden;height:0;position:relative;\" class=\"uk-hidden\"><ul class=\"uk-nav-sub\" role=\"menu\">\n          <li><a href=\"dropdown.html\">Dropdown</a></li>\n          <li><a href=\"modal.html\">Modal</a></li>\n          <li><a href=\"offcanvas.html\">Off-canvas</a></li>\n          <li><a href=\"switcher.html\">Switcher</a></li>\n          <li><a href=\"toggle.html\">Toggle</a></li>\n          <li><a href=\"scrollspy.html\">Scrollspy</a></li>\n          <li><a href=\"smooth-scroll.html\">Smooth scroll</a></li>\n        </ul></div>\n      </li>\n      <li class=\"uk-nav-header\">Components</li>\n      <li class=\"uk-parent\" aria-expanded=\"false\"><a href=\"#\"><i class=\"uk-icon-th-large\"></i> Layout</a>\n        <div style=\"overflow:hidden;height:0;position:relative;\" class=\"uk-hidden\"><ul class=\"uk-nav-sub\" role=\"menu\">\n          <li><a href=\"grid-js.html\">Dynamic Grid</a></li>\n          <li><a href=\"grid-parallax.html\">Parallax Grid</a></li>\n        </ul></div>\n      </li>\n      <li class=\"uk-parent\" aria-expanded=\"false\"><a href=\"#\"><i class=\"uk-icon-bars\"></i> Navigations</a>\n        <div style=\"overflow:hidden;height:0;position:relative;\" class=\"uk-hidden\"><ul class=\"uk-nav-sub\" role=\"menu\">\n          <li><a href=\"dotnav.html\">Dotnav</a></li>\n          <li><a href=\"slidenav.html\">Slidenav</a></li>\n          <li><a href=\"pagination-js.html\">Dynamic Pagination</a></li>\n        </ul></div>\n      </li>\n      <li class=\"uk-parent\" aria-expanded=\"false\"><a href=\"#\"><i class=\"uk-icon-folder-open\"></i> Common</a>\n        <div style=\"overflow:hidden;height:0;position:relative;\" class=\"uk-hidden\"><ul class=\"uk-nav-sub\" role=\"menu\">\n          <li><a href=\"form-advanced.html\">Form advanced</a></li>\n          <li><a href=\"form-file.html\">Form file</a></li>\n          <li><a href=\"form-password.html\">Form password</a></li>\n          <li><a href=\"form-select.html\">Form select</a></li>\n          <li><a href=\"placeholder.html\">Placeholder</a></li>\n          <li><a href=\"progress.html\">Progress</a></li>\n        </ul></div>\n      </li>\n      <li class=\"uk-parent\" aria-expanded=\"false\"><a href=\"#\"><i class=\"uk-icon-magic\"></i> JavaScript</a>\n        <div style=\"overflow:hidden;height:0;position:relative;\" class=\"uk-hidden\"><ul class=\"uk-nav-sub\" role=\"menu\">\n          <li><a href=\"lightbox.html\">Lightbox</a></li>\n          <li><a href=\"autocomplete.html\">Autocomplete</a></li>\n          <li><a href=\"datepicker.html\">Datepicker</a></li>\n          <li><a href=\"htmleditor.html\">HTML editor</a></li>\n          <li><a href=\"slider.html\">Slider</a></li>\n          <li><a href=\"slideset.html\">Slideset</a></li>\n          <li><a href=\"slideshow.html\">Slideshow</a></li>\n          <li><a href=\"parallax.html\">Parallax</a></li>\n          <li><a href=\"accordion.html\">Accordion</a></li>\n          <li><a href=\"notify.html\">Notify</a></li>\n          <li><a href=\"search.html\">Search</a></li>\n          <li><a href=\"nestable.html\">Nestable</a></li>\n          <li><a href=\"sortable.html\">Sortable</a></li>\n          <li><a href=\"sticky.html\">Sticky</a></li>\n          <li><a href=\"timepicker.html\">Timepicker</a></li>\n          <li><a href=\"tooltip.html\">Tooltip</a></li>\n          <li><a href=\"upload.html\">Upload</a></li>\n        </ul></div>\n      </li>\n      <li class=\"uk-nav-divider\"></li>\n      <li><a href=\"../showcase/index.html\">Showcase</a></li>\n      <li><a href=\"tutorials.html\">Tutorials</a></li>\n      <li><a href=\"/\">UIkit 3</a></li>\n    </ul>\n\n  </div>\n\n</div>\n")
	return nil
}
