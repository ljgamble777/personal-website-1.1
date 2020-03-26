(function() {
    // Get domain so we don't attach this listener to internal links.
    var thisDomain = document.domain.split('.').reverse()[1] + '.' + document.domain.split('.').reverse()[0];

    // Get list of all a links.
    var aList = document.getElementsByTagName('a');

    for (i = 0; i < aList.length; i++) {
        a = aList[i]
        // Attach a listener to figure the gtag marking the event.
        if (!a.getAttribute("href").includes(thisDomain)) {
            a.addEventListener("click", function(event) {
                gtag('event', 'external_link_click', {
                    'event_category': 'external_link',
                    'event_label': event.target.getAttribute("data-label")
                });
            });
        }
    }
})();