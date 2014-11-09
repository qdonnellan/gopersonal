// This is a simple *viewmodel* - JavaScript that defines the data and behavior of your UI
function EssayViewModel() {
    var self = this;
    self.rawMarkdown = ko.observable('');
    self.parsedMarkdown = ko.computed(function() {
        return markdown.toHTML(self.rawMarkdown());
    }); 
};

$(document).ready(function() {
    // initialte the view model
    var vm = new EssayViewModel()
    ko.applyBindings(vm);

    // parse the raw markdown data which is embedded in the dom,
    // put there by the server when rendering the essay template
    var raw_markdown = $(".raw-essay-markdown").html()
    vm.rawMarkdown(raw_markdownm);
});
