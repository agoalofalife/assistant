
const websocket = {
    install(Vue, options) {
        Vue.mixin({
            mounted() {
                // new Websocket("ws://localhost:7777/echo");
                console.log( Websocket );
                console.log('Mounted!');
            }
        });
    }
};

export default websocket;