// This is a simple *viewmodel* - JavaScript that defines the data and behavior of your UI
function FrontPageViewModel() {
    var self = this;
    self.formEmail = ko.observable('');
    self.formEmailError = ko.observable(false);
    self.formContent = ko.observable('');
    self.successfulSubmission = ko.observable(false);
    self.submissionErrorMsg = ko.observable(null);
    self.isFormEmailValid = ko.computed(function() {
        var validityStatus = true;
        if ( self.formEmail() == '' ) {
            validityStatus = false;
        }
        return validityStatus
    });

    self.validateAndSubmitContactForm = function() {
        if ( self.isFormEmailValid() ) {
            self.submitContactForm();
        } else {
            self.formEmailError(true)
        }
    }

    self.submitContactForm = function() {
        var data = {
            email: self.formEmail(),
            content: self.formContent()
        }
        $.post("/contact", data, function(response_data) {
            if (response_data.success == true) {
                self.successfulSubmission(true);
            } else {
                self.submissionErrorMsg('error');
            }
        });
    }
    
};

$(document).ready(function() {
    // initialte the view model
    var vm = new FrontPageViewModel()
    ko.applyBindings(vm);
});
