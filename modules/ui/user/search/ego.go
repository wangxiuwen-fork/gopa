// Generated by ego.
// DO NOT EDIT

package search

import (
	"fmt"
	"github.com/medcl/gopa/modules/ui/admin/common"
	"io"
)

var _ = fmt.Sprint("") // just so that we can keep the fmt import for now
func Search(w io.Writer) error {
	_, _ = io.WriteString(w, "\n\n")
	_, _ = io.WriteString(w, "\n\n")
	common.Head(w, "Search", "")
	_, _ = io.WriteString(w, "\n")
	common.Body(w)
	_, _ = io.WriteString(w, "\n")
	common.Nav(w, "Search")
	_, _ = io.WriteString(w, "\n\n\n<META HTTP-EQUIV=\"refresh\" CONTENT=\"0;URL=/admin/\">\n\n<div class=\"tm-middle\">\n    <div class=\"uk-container uk-container-center\">\n\n        <div class=\"uk-grid\" data-uk-grid-margin=\"\">\n            <div class=\"tm-sidebar uk-width-medium-1-4 uk-hidden-small uk-row-first\">\n\n\n\n            </div>\n            <div class=\"tm-main uk-width-medium-3-4\">\n\n                <article class=\"uk-article\">\n\n                    Welcome to GOPA.\n\n                </article>\n\n            </div>\n        </div>\n\n    </div>\n</div>\n\n")
	common.OffCanvas(w)
	_, _ = io.WriteString(w, "\n\n")
	common.Footer(w)
	_, _ = io.WriteString(w, "\n")
	return nil
}
