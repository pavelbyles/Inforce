$(function() {
	
	activateSelectCompany();

	$('#gopher').each(function(){
		    $(this).height($(this).height() * 0.05);
	});

	$('#addcompanyform').submit(function(event) {
		var f = $(this);
		$.post(f.attr('action'), f.serialize(), function(data) {
			doAlert("Successfully added company", "#addcompanyresult");
			getCompanies();
			$("#addcompanyform").trigger("reset");
		});
		event.preventDefault();
	});
});

function activateSetUser() {
	$('#setUserForm').submit(function(event) {
		var f = $(this);
		$.post(f.attr('action'), f.serialize(), function(data) {
			$("#selectCompanyForm").submit();
		});
		event.preventDefault();
	});
}

function activateAddUserToCompany() {
	$('#addusertocompany').submit(function(event) {
		var f = $(this);
		$.post(f.attr('action'), f.serialize(), function(data) {
			$('#selectCompanyForm').submit();
		});
		event.preventDefault();
	});
}

function activateSendMsg() {
	$('#sendcompmsg').submit(function(event) {
		var f = $(this);
		$.post(f.attr('action'), f.serialize(), function(data) {
			doAlert("Message successfully sent", "#msgSendResult");
			$('#selectCompanyForm').submit();
		});
		event.preventDefault();
	});
}

function doAlert(message, targetDiv) {
	$(targetDiv).html(message);
	$(targetDiv).addClass("alert");
	$(targetDiv).css("display", "block");
	$(targetDiv + ".alert").alert();
	setTimeout(function() {
		$(targetDiv).css("display", "none");
		$(targetDiv).removeClass("alert");
	}, 3000);
}

function getCompanies() {
	$.post('getCompanies?type=html', function(data) {
		$('#companiesSection').html(data);
		activateSelectCompany();
	});
}

function activateSelectCompany() {
	$('#selectCompanyForm').submit(function(event) {
		var f = $('#selectCompanyForm');
		$.post(f.attr('action'), f.serialize(), function(data) {
			$('#companyInfo').html(data);
			activateSelectDept();
			activateAddUserToCompany();
			activateSetUser();
			activateSendMsg();
		});
		event.preventDefault();
	});
	$(".dropdown-toggle").dropdown();
}

function activateSelectDept() {
	$('#adddepartmentform').submit(function(event) {
		var f = $(this);
		$.post(f.attr('action'), f.serialize(), function(data) {
			$("#selectCompanyForm").submit();
		})
		event.preventDefault();
	});
}

function getDepartments() {
	$.post('getDepts?type=html', function(data) {
		$('#departmentsSection').html(data);
		activateSelectDept();
	});
}

function addCompanyUser() {
	var f = $(this);
	$.post(f.attr('action'), f.serialize(), function(data) {
		activateSelectCompany();
	});
}

