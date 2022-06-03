jQuery.each(["put", "delete"], function (i, method) {
    jQuery[method] = function (url, data, callback, type) {
        if (jQuery.isFunction(data)) {
            type = type || callback;
            callback = data;
            data = undefined;
        }
        return jQuery.ajax({
            url: url,
            type: method,
            dataType: 'json',
            data: data,
            success: callback
        });
    };
});
//https://codeseven.github.io/toastr/demo.html
toastr.options = {
    "closeButton": true,
    "debug": false,
    "newestOnTop": false,
    "progressBar": true,
    "positionClass": "toast-top-center",
    "preventDuplicates": false,
    "onclick": null,
    "showDuration": "300",
    "hideDuration": "1000",
    "timeOut": "5000",
    "extendedTimeOut": "1000",
    "showEasing": "swing",
    "hideEasing": "linear",
    "showMethod": "fadeIn",
    "hideMethod": "fadeOut"
}

function noticeOk(msg) {
    toastr.options.timeOut = 2000;
    toastr['success'](msg);
}

function noticeError(msg) {
    toastr.options.timeOut = 5000;
    toastr['error'](msg);
}

function noticeWarning(msg) {
    toastr.options.timeOut = 5000;
    toastr['warning'](msg);
}

function noticeInfo(msg) {
    toastr.options.timeOut = 5000;
    toastr['info'](msg);
}

function logout() {
    $.ajax({
        url: '/admin/logout', type: 'get', dataType: 'json', success: function (data) {
            console.log(data)
            if (data.code == 0) {
                window.location.href = '/login';
            }
        }
    });
}


function updatePwd() {
    let old = prompt('Input your old pwd');
    if (!old) {
        return
    }
    let newPwd = prompt('Input your new pwd');
    if (!newPwd) {
        return
    }
    $.post("/admin/updatePwd", {oldPwd: old, newPwd: newPwd}, (res) => {
        if (res.code == 0) {
            noticeOk('Success')
            location.href = '/login'
        } else {
            noticeError(res.msg)
        }
    });
}

// 添加默认选中
$(".sub-nav a").hide()
let current = $("a[href='" + location.pathname + "']");
$(".sub-nav a[data='" + current.attr("data") + "']").show()
$(".nav a[data='" + current.attr("data") + "']").addClass("link-2-active")


// 监听tab 切换
$(function () {
    $(".nav a").hover(function () {
        $(".nav a").removeClass("link-2-active")
        $(this).addClass("link-2-active")
        $(".sub-nav a").hide()
        $(".sub-nav a[data='" + $(this).attr("data") + "']").show()
    })
})

function setDark(value) {
    Cookies.set('dark', value)
    location.reload()
}