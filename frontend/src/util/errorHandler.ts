
// 401認証エラー時の処理
export const authErrorMessage = () => {
    alert("セッションが切れました。もう一度ログインしてください。");
    location.reload();
};
// 500サーバーエラー時の処理
export const serverErrorMessage = () => {
    alert("処理が正常に行われませんでした。");
    location.reload();
};