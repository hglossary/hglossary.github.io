import './main.scss';
import Main from './Main.svelte';
import {setupStyle} from './share/inline-images.js';

setupStyle();

export default function main() {
  const app = new Main({
    target: document.body,
  });
}

main();
