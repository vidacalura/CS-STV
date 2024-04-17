//const API = "http://127.0.0.1:4000/api";
const API = "https://cs-stv.onrender.com/api";

const mostrarDadosGeraisDupla = async (fetchFunc, path) => {
    const res = await fetchFunc(path);
    mostrarMsgErro(res.error);

    globalThis.dupla = Object.freeze(res.dupla);
    const dupla = globalThis.dupla;

    const logo = document.getElementById("img-dupla");
    atualizarImagemSrc(logo, concatStr("imgs/", dupla.logo));

    const nome = document.getElementById("nome-dupla");
    atualizarTexto(nome, dupla.nome);

    const rank = document.getElementById("ranking-dupla");
    atualizarTexto(rank, concatStr("#", dupla.rank));
};

/* Roster */
const mostrarInfoRoster = async(fetchFunc, path) => {
    const res = await fetchFunc(path);
    mostrarMsgErro(res.error);

    const roster = res.dupla.roster;

    const template = `
        <h3> Jogadores de Dupla </h3>
        <div class="tabela-dupla-roster">
            <div class="tabela-dupla-roster-legenda">
                <p> Jogador </p>
                <p> Mapas jogados </p>
                <p> Função </p>
            </div>

            <div class="tabela-dupla-roster-coluna">
                <p> ${roster.IGL} </p>
                <p> 20.000 </p>
                <p> IGL </p>
            </div>
        </div>`;
};

/* Partidas */
const mostrarInfoPartidas = async(fetchFunc, path) => {
    const res = await fetchFunc(path);
    mostrarMsgErro(res.error);

    const infoContainer = document.getElementById("container-info-dupla");
    atualizarInnerHTML(infoContainer,
        concatStr("<h3>Partidas de ",
            concatStr(globalThis.dupla.nome, "</h3>")));

    const prtdArr = res.partidas;
    prtdArr.map(appendPartidaInfoContainer);
};

const appendPartidaInfoContainer = prtd => {
    const templatePrtdHTML = `
        <div class="container-partida-dupla">
            <p class="data-partida-dupla">
                ${formatarData(prtd.dataJogo)}
            </p>

            <div class="container-resultados-partida-dupla">
                <div>
                    <img src="imgs/${prtd.timeCasa.logoURL || "tr.png"}" 
                        alt="Logo dupla 1">
                    <p class="nome-resultados-partida-dupla"> ${prtd.timeCasa.nome} </p>
                    <p class="${
                        prtd.timeCasa.pontos + prtd.timeFora.pontos >= 16
                        ? prtd.timeCasa.pontos > prtd.timeFora.pontos
                        ? "res-partida-vitoria"
                        : prtd.timeCasa.pontos < prtd.timeFora.pontos
                        ? "res-partida-derrota"
                        : null
                        : null
                    }"> 
                        ${prtd.timeCasa.pontos}
                    </p>
                </div>
                <p> : </p>
                <div>
                    <p class="${
                        prtd.timeCasa.pontos + prtd.timeFora.pontos >= 16
                        ? prtd.timeCasa.pontos < prtd.timeFora.pontos
                        ? "res-partida-vitoria"
                        : prtd.timeCasa.pontos > prtd.timeFora.pontos
                        ? "res-partida-derrota"
                        : null
                        : null
                    }"> ${prtd.timeFora.pontos} </p>
                    <p class="nome-resultados-partida-dupla"> ${prtd.timeFora.nome} </p>
                    <img src="imgs/${prtd.timeFora.logoURL || "tr.png"}" 
                        alt="Logo dupla 2">
                </div>
            </div>
        </div>`;

    const infoContainer = document.getElementById("container-info-dupla");
    infoContainer.innerHTML = concatStr(infoContainer.innerHTML, templatePrtdHTML);
}


mostrarDadosGeraisDupla(fetchCSSTVAPI, 
    concatStr("/duplas/", getURLParam("codDupla")));

document.getElementById("dupla-partidas-button").addEventListener("click", () =>
    mostrarInfoPartidas(fetchCSSTVAPI,
        concatStr("/partidas/dupla/", getURLParam("codDupla"))));