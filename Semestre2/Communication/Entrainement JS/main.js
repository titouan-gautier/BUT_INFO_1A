/* const title = document.querySelector('h1');

window.addEventListener( 'scroll' , () => {

    const {scrollTop,clientHeight} = document.documentElement;

    console.log(scrollTop,clientHeight)

    if (scrollTop == 0) {
        title.classList.add('.color-title')
    }

 })
 */
 const slidingNewsletter = document.querySelector('.slide-in');

window.addEventListener('scroll', () => {

    const {scrollTop, clientHeight} = document.documentElement;

    console.log(scrollTop, clientHeight);
})