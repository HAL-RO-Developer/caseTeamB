import axios from 'axios'

class Http {

    constructor() {
        this.Load()
    }

    Load() {
        this.api = axios.create({
            baseURL: 'http://localhost:8080',
            //baseURL: 'https://caseteamaserver-lyozxwqhak.now.sh',
            //baseURL: localStorage.getItem('server'),
            headers: {
                'ContentType': 'application/json',
                'Authorization': this.GetToken()
            },
            responseType: 'json'
        })
    }

    SetToken(t) {
        localStorage.setItem("token", t)
        http.Load();
    }
    GetToken() {
        return localStorage.getItem("token")
    }
    RemoveToken() {
        localStorage.removeItem("token")
    }
    // サインイン
    signin(name, pass) {
        return this.api.post('/signin', {
            name,
            pass
        })
    }
    // サインアップ
    signup(name, pass) {
        return this.api.post('/signup', {
            name,
            pass
        })
    }
    // ユーザー削除
    removeUser() {
        return this.api.delete('/user')
    }
    // デバイスID発行
    registDevice(goal_id) {
        return this.api.post('/device', {
            goal_id
        })
    }
    // デバイスID取得
    getDevice() {
        return this.api.get('/device')
    }
    // デバイスID削除
    removeDevice(device_id) {
        return this.api.delete('/device/' + device_id)
    }
    // 子ども情報取得
    getChild() {
        return this.api.get('/child')
    }
    // 子ども情報追加
    addChild(nickname, birthday, sex) {
        return this.api.post('/child', {
            nickname,
            birthday,
            sex
        })
    }
    // 子ども情報削除
    removeChild(child_id) {
        return this.api.delete('/child/' + child_id)
    }
    // 目標追加
    addGoal(child_id, content, criteria, deadline) {
        return this.api.post('/goal', {
            child_id,
            content,
            criteria,
            deadline
        })
    }
    // 目標取得
    getGoal() {
        return this.api.get('/goal')
    }
    // 目標削除
    removeGoal(goal_id) {
        return this.api.delete('/goal/' + goal_id, {
        })
    }
    // 達成数変更
    putAchieved(goal_id, approval) {
        return this.api.put('/approval', {
            goal_id,
            approval
        })
    }
    // メッセージ追加
    addMessage(goal_id, message_call, message) {
        return this.api.post('/message', {
            goal_id,
            message_call,
            message
        })
    }
    // メッセージ取得
    getMessage() {
        return this.api.get('/message')
    }
    // メッセージ削除
    removeMessage(goal_id, message_call) {
        return this.api.delete('/message/' + goal_id +'/'+ message_call)
    }
}
var http = new Http()
export default http;
