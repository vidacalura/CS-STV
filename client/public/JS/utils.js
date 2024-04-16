const atualizarTexto = (DOMElement, str) => DOMElement.textContent = str;
const atualizarInnerHTML = (DOMElement, str) => DOMElement.innerHTML = str;
const atualizarImagemSrc = (DOMElement, src) => DOMElement.src = src;
const atualizarHref = (DOMElement, link) => DOMElement.href = link;
const concatStr = (str1, str2) => str1.concat(str2);
const getURLParam = (param) => 
    new URLSearchParams(window.location.search).get(param);
const mostrarMsgErro = err => err != null ? alert(err) : null;

const fetchCSSTVAPI = async pathAPI => 
    fetch(concatStr(API, pathAPI))
        .then(res => res.json())
        .then(res => res);

const formatarData = data => 
    concatStr(data.split("T")[0].split("-").reverse().join("/"), 
        concatStr(" ", data.split("T")[1].slice(0, -4)));