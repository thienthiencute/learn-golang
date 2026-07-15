import Alert from "/static/js/components/alert.js";
import { getCookie, saveAuth } from "/static/js/common/helpers.js";

const API_BASE_URL = document.querySelector('meta[name="api-base-url"]')?.getAttribute('content') || "";

function handleErrorResponse(responseData) {
    if (responseData?.details?.msg && Array.isArray(responseData.details.msg)) {
        return responseData.details.msg.join('\n');
    }

    if (responseData?.message) {
        return responseData.message;
    }

    return 'Đã xảy ra lỗi không xác định.';
}

async function login(urlLogin, email, password) {
    try {
        const response = await fetch(urlLogin, {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
                'X-CSRF-Token': getCookie('csrf_'),
            },
            body: JSON.stringify({ email, password }),
        });

        if (!response.ok) {
            const errorData = await response.json().catch(() => null);
            const errorMessage = handleErrorResponse(errorData);
            Alert.error(errorMessage);
            return;
        }
       

        const datas = await response.json();

        if (datas?.data?.access_token?.token) {
            handleLoginSuccess(datas.data)
        } else {
            Alert.error('Không nhận được access token từ server.');
        }
    } catch (error) {      
        Alert.error('Có lỗi xảy ra khi đăng nhập. Vui lòng thử lại.');
    }
}

function handleLoginSuccess(data) {   
    saveAuth(data.user); 
    const urlParams = new URLSearchParams(window.location.search);
    const redirectUrl = urlParams.get('redirect') || '/';
    window.location.href = redirectUrl;
}

async function signup(urlSignup, data) {
    try {
        const response = await fetch(urlSignup, {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
                'X-CSRF-Token': getCookie('csrf_'),
            },
            body: JSON.stringify(data),
        });

        const responseText = await response.text();

        let responseBody;
        try {
            responseBody = JSON.parse(responseText);
        } catch (e) {
            responseBody = responseText;
        }

        if (response.ok) {
            Alert.success("Registered successfully!");
            setTimeout(() => {
                window.location.href = "/login";
            }, 1200);
            return;
        }

        if (responseBody && typeof responseBody === 'object' && responseBody.details && responseBody.details.msg) {
            Alert.error(responseBody.details.msg);
        } else {
            const errorMessage = handleErrorResponse(responseBody);
            Alert.error(errorMessage);
        }

    } catch (error) {      
        Alert.error('There was an error during registration. Please try again.');
    }
}

function getHeaders() {
    return {
        'Content-Type': 'application/json',
        "X-CSRF-Token": getCookie("csrf_")
    };
}

function logout() {   
    document.cookie = "brx_token=; expires=Thu, 01 Jan 1970 00:00:00 UTC; path=/; Secure; SameSite=Strict";
    localStorage.removeItem('access_token');
    localStorage.removeItem('access_token_exp');
    localStorage.removeItem('user_info');
    window.location.href = "/login";
}

export { 
    login,   
    signup,
    handleErrorResponse,
    getHeaders,
};

