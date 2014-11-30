function EssayListViewModel() {
    var self = this;
    self.essays = ko.observableArray([
        { essayRef : "working-from-home-with-a-newborn-the-first-three-months" },
    ]);

    self.formattedEssays = ko.computed(function() {
        var formattedArray = [];
        var i, essay, arrayItem;
        for ( i in self.essays() ) {
            essay = self.essays()[i];
            arrayItem = {
                title: essay.essayRef.replace(/-/g, " "),
                link: "/essays/" + essay.essayRef
            }
            formattedArray.push(arrayItem);
        }
        return formattedArray;
    });
};

$(document).ready(function() {
    var vm = new EssayListViewModel();
    ko.applyBindings(vm);
});

