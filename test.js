console.log("load ok");

setInterval(function() {

    var array = [0x11, 0x02, 0x03, 0x04, 0x05, 0x06, 0x27];
    var arrayBuffer = new Uint8Array(array).buffer;

    send({"dd":"dd"},array);

},1000);
