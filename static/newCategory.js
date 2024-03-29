let leftMouseClick = false;

document.addEventListener('mousedown', function(e) {
    leftMouseClick = e.button === 0;
})
document.addEventListener('mouseup', function() {
    leftMouseClick = false;
})

function check(id) {
    if (leftMouseClick) {
        let checkbox = document.getElementById(id);
        checkbox.click();
    }
}

async function submit() {
    let mangas = []
    document.querySelectorAll(".list").forEach(function(element) {
        if (element.checked) {
            mangas.push(element.id)
        }
    })

    let name = document.getElementById("category-name").value;

    if (name === "") {
        alert("Please enter a name for the category");
        return;
    }

    let url = window.location.href;
    await fetch(url, {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify({
            name: name,
            mangas: mangas,
        })
    }).then(function(response) {
        if (response.status === 200) {
            alert("Category created!");
            document.getElementById("category-name").value = "";
            document.querySelectorAll(".list").forEach(function(element) {
                element.checked = false;
            })
        } else {
            alert("Error creating category!\n" + response.body);
        }
    })
}