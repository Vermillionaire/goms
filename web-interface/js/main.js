$( document ).ready(function() {
    console.log("Loading html...");
    $("#media-cards").load("/media/")
    console.log("Done loading html");
    
});


var $grid = $('.grid').masonry({
  // options
  itemSelector: '.grid-item',
  columnWidth: '.grid-sizer',
  transitionDuration: '0.2s',
  initLayout: false
});


$grid.imagesLoaded().progress( function() {
    console.log("Done loading images...");
    
    $grid.masonry();
});