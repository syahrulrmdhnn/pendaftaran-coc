window.onload = function(){
  const apiUrl = `http://localhost:5000/api/get/coc02/coconut@013`;
  fetch(apiUrl)
  .then(response => {
    if (!response.ok){
      throw new Error("Gagal mengambil data");
    }
    return response.json();
  })
  .then(data => {
    if (data.message === "success") {
      console.log("Data berhasil dimuat:", data);
      // Process the data as needed
      // The data structure is: {"message":"success","data":{"id":{"nama_lengkap":"...", "email":"...", ...}}}
    } else {
      console.log("Tidak ada data pendaftar");
    }
  })
  .catch(error => {
    console.error("Terjadi kesalahan:", error);
  })
}