var data = []

function slugify(text) {
    return text.toString().toLowerCase()
        .replace(/\s+/g, '-')           // Replace spaces with -
        .replace(/[^\w\-]+/g, '')       // Remove all non-word chars
        .replace(/\-\-+/g, '-')         // Replace multiple - with single -
        .replace(/^-+/, '')             // Trim - from start of text
        .replace(/-+$/, '');            // Trim - from end of text
}

document.querySelector("#content > div.container.pb-5 > div > div.col-lg-auto.category-tree-col > div > ul").childNodes.forEach(e => {
    let children = []
    if (e.childNodes.length > 1) {
        e.childNodes[1].childNodes.forEach(x => {
            children.push({
                name: "",
                en_name: e.children[0].innerText,
                slug: slugify(e.children[0].innerText),
                children: []
            })
        })
    }

    data.push({
        name: "",
        en_name: e.children[0].innerText,
        slug: slugify(e.children[0].innerText),
        children: children
    })

})