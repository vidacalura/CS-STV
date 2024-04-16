//const API = "http://127.0.0.1:4000/api";
const API = "https://cs-stv.onrender.com/api";

const mostrarNoticia = async (fetchFunc, path) => {
    const res = await fetchFunc(path);
    mostrarMsgErro(res.error);

    const notc = res.noticia;
    
    const titulo = document.getElementById("noticia-titulo");
    atualizarTexto(titulo, notc.titulo);

    const subtitulo = document.getElementById("noticia-subtitulo");
    atualizarTexto(subtitulo, notc.subtitulo);

    const data = document.getElementById("data-noticia");
    atualizarTexto(data, formatarData(notc.data));

    const corpoNoticia = document.getElementById("corpo-noticia");
    atualizarInnerHTML(corpoNoticia, notc.noticia);
};

mostrarNoticia(fetchCSSTVAPI, concatStr("/noticias/", getURLParam("codNotc")));