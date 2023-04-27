const goWasm = new Go()

WebAssembly.instantiateStreaming(fetch("main.wasm"), goWasm.importObject)
    .then((result) => {
        goWasm.run(result.instance)

        document.getElementById("left").addEventListener("click", () => {
            left()
        })

        document.getElementById("right").addEventListener("click", () => {
            right()
        })

        document.getElementById("up").addEventListener("click", () => {
            up()
        })

        document.getElementById("down").addEventListener("click", () => {
            down()
        })
    })

