//const API = "http://127.0.0.1:4000/api";
const API = "https://cs-stv.onrender.com/api";

/* Player of the week */
const mostrarPlayerOfTheWeek = async (fetchFunc, path) => {
    const res = await fetchFunc(path);
    const playerData = res.playerData;
    // checar erro

    const nameContainer = document.getElementById("player-of-the-week-name");
    const dadosContainer = document.getElementById("player-of-the-week-dados");
    const infoContainer = document.getElementById("player-of-the-week-info-dados");

    atualizarTexto(nameContainer, playerData.nome);
    atualizarTexto(dadosContainer, playerData.dados);
    atualizarTexto(infoContainer, playerData.infoDados);
};

/* Ranking Duplas */
const criarElementoRankingDuplas = dupla => {
    const linkContainer = document.createElement("a");
    atualizarHref(linkContainer, concatStr("dupla.html?codDupla=", dupla.codDupla));

    const containerDupla = document.createElement("div");
    containerDupla.classList.add("dupla-ranking-container");

    const lugarRanking = document.createElement("p");
    atualizarTexto(lugarRanking, concatStr(dupla.rank.toString(), "."));

    const logoDupla = document.createElement("img");
    atualizarImagemSrc(logoDupla, concatStr("./imgs/", dupla.logo || "tr.png"));

    const nomeDupla = document.createElement("p");
    atualizarTexto(nomeDupla, dupla.nome);

    containerDupla.appendChild(lugarRanking);
    containerDupla.appendChild(logoDupla);
    containerDupla.appendChild(nomeDupla);

    linkContainer.appendChild(containerDupla);

    const container = document.getElementById("container-top-5-duplas-home");
    container.appendChild(linkContainer);
};

const mostrarRankingDuplas = async (fetchFunc, path) => {
    const res = await fetchFunc(path);
    mostrarMsgErro(res.error);

    const rankingArr = res.ranking;
    rankingArr.map(criarElementoRankingDuplas);
};

/* Partidas recentes */
const getClasseResultadoPartida = (pontosTimeA, pontosTimeB) => 
    pontosTimeA > pontosTimeB 
    ? "res-partida-vitoria"
    : pontosTimeA === pontosTimeB 
    ? "res-partida-empate"
    : "res-partida-derrota";

const criarElementoPartida = prtd => {
    const linkContainer = document.createElement("a");
    atualizarHref(linkContainer, concatStr("partida.html?codPrtd=", prtd.codPrtd));

    const prtdContainer = document.createElement("div");
    prtdContainer.classList.add("partida-home-container");

    const dataPartida = document.createElement("p");
    dataPartida.classList.add("data-partida-home");
    atualizarTexto(dataPartida, formatarData(prtd.dataJogo));
    prtdContainer.appendChild(dataPartida);

    const timeCasaContainer = document.createElement("div");
    timeCasaContainer.classList.add("partida-home-dupla-container");

    const timeCasaNome = document.createElement("p");
    atualizarTexto(timeCasaNome, prtd.timeCasa.nome);
    timeCasaContainer.appendChild(timeCasaNome);

    const timeCasaPontos = document.createElement("p");
    timeCasaPontos.classList.add(
        getClasseResultadoPartida(prtd.timeCasa.pontos, prtd.timeFora.pontos));
    atualizarTexto(timeCasaPontos, concatStr("(",
        concatStr(prtd.timeCasa.pontos.toString(), ")")));
    timeCasaContainer.appendChild(timeCasaPontos);

    const timeForaContainer = document.createElement("div");
    timeForaContainer.classList.add("partida-home-dupla-container");

    const timeForaNome = document.createElement("p");
    atualizarTexto(timeForaNome, prtd.timeFora.nome);
    timeForaContainer.appendChild(timeForaNome);

    const timeForaPontos = document.createElement("p");
    timeForaPontos.classList.add(
        getClasseResultadoPartida(prtd.timeFora.pontos, prtd.timeCasa.pontos));
    atualizarTexto(timeForaPontos,
        concatStr("(", concatStr(prtd.timeFora.pontos.toString(), ")")));
    timeForaContainer.appendChild(timeForaPontos);

    prtdContainer.appendChild(timeCasaContainer);
    prtdContainer.appendChild(timeForaContainer);

    linkContainer.appendChild(prtdContainer);

    const container = document.getElementById("container-partidas-home");
    container.appendChild(linkContainer);
};

const mostrarPartidasRecentes = async (fetchFunc, path) => {
    const res = await fetchFunc(path);
    mostrarMsgErro(res.error);

    const partidasArr = res.partidas;
    partidasArr.map(criarElementoPartida);
};

/* Eventos recentes */
const criarElementoEvento = evnt => {
    const linkContainer = document.createElement("a");
    atualizarHref(linkContainer, concatStr("evento.html?codEvnt=", evnt.codEvnt));

    const eventoContainer = document.createElement("div");
    eventoContainer.classList.add("evento-home-container");

    const nomeEvento = document.createElement("p");
    atualizarTexto(nomeEvento, evnt.evento);
    eventoContainer.appendChild(nomeEvento);

    linkContainer.appendChild(eventoContainer);

    const container = document.getElementById("container-eventos-home");
    container.append(linkContainer);
};

const mostrarEventosRecentes = async (fetchFunc, path) => {
    const res = await fetchFunc(path);
    mostrarMsgErro(res.error);

    const eventosArr = res.eventos;
    eventosArr.map(criarElementoEvento);
}

/* Feed de notícias */
const criarElementoNoticia = notc => {
    const linkContainer = document.createElement("a");
    atualizarHref(linkContainer, concatStr("noticia.html?codNotc=", notc.codNotc));

    const noticiaContainer = document.createElement("div");
    noticiaContainer.classList.add("container-noticias-home");

    const dataNoticia = document.createElement("p");
    atualizarTexto(dataNoticia, formatarData(notc.data));
    noticiaContainer.appendChild(dataNoticia);

    const tituloNoticia = document.createElement("h3");
    atualizarTexto(tituloNoticia, notc.titulo);
    noticiaContainer.appendChild(tituloNoticia);

    linkContainer.appendChild(noticiaContainer)

    const container = document.getElementById("container-noticias-home");
    container.appendChild(linkContainer);
};

const mostrarFeedNoticias = async (fetchFunc, path) => {
    const res = await fetchFunc(path);
    mostrarMsgErro(res.error);

    const noticiasArr = res.feed;
    noticiasArr.map(criarElementoNoticia);
};

mostrarPlayerOfTheWeek(fetchCSSTVAPI, "/player-of-the-week");
mostrarRankingDuplas(fetchCSSTVAPI, "/duplas/ranking");
mostrarPartidasRecentes(fetchCSSTVAPI, "/partidas/recentes");
mostrarEventosRecentes(fetchCSSTVAPI, "/eventos/recentes");
mostrarFeedNoticias(fetchCSSTVAPI, "/noticias/feed");