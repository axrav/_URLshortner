var inp = document.getElementById("url");
    inp.addEventListener("keypress", function(event){
        if (event.key === "Enter") {
            event.preventDefault();
            document.getElementById("submitbutton").click();
    }
    })
    function isUrl(url) {
        let regEx = /^https?:\/\/(?:www\.)?[-a-zA-Z0-9@:%._\+~#=]{1,256}\.[a-zA-Z0-9()]{1,6}\b([-a-zA-Z0-9()@:%_\+.~#?&//=]*)$/gm;
        return regEx.test(url);
        }
    function shortit(){
        var r = document.getElementById("url").value;
        const withHttp = url => !/^https?:\/\//i.test(url) ? `http://${url}` : url;
        r = withHttp(r)
        if (r.length === 0){
            alert("URL cant be empty")
            return
        }
        if (!isUrl(r)){
            alert("You Entered an invalid URL")
            return
        }
        fetch("/short",
        {
            headers: {
                'Accept': 'application/json',
                'Content-Type': 'application/json;charset=UTF-8'
                    },
            method: "POST",
            body: JSON.stringify({"url" : r})
}).then(res=> res.json()).then(function (data) { var d = JSON.parse(data)
    document.getElementById("shortened").innerHTML = d.shortened_url
    document.getElementById("shortened").setAttribute("href", d.shortened_url);
})
    }