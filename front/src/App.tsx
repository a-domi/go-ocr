import axios from "axios";
import React from 'react';
import { useEffect,useState,ChangeEvent } from 'react';
import './App.css';

function App() {
  const [file, setFile] = useState<File | null>(null);
  const [inProcess, setInProcess] = useState(false);
  const [previewFileUrl, setPreviewFileUrl] = useState('');
  const [resultText, setResultText] = useState('');

  useEffect(()=>{
    
  })

  const onChangeFile = (e: ChangeEvent<HTMLInputElement>) => {
    setResultText('');
    if(!(e.target instanceof HTMLInputElement)) return
    if(!e.target.files) return
    const file = e.target.files[0];
    setFile(file);
    const previewUrl = file ? URL.createObjectURL(file) : '';
    setPreviewFileUrl(previewUrl)
  }

  const onClickUpload = async() => {
    setInProcess(true);
    if (file) {
      const data = new FormData();
      data.append('file', file);
      axios.post('http://localhost:8080',data)
      .then((res) => {
        console.log(res.data);
        setResultText(res.data);
        setInProcess(false);
      })
      .catch((err) => {
        alert('error');
        setPreviewFileUrl('');
        setInProcess(false);    
      })
    }
  }

  return (
    <div  className="main-container">
      <div className="upload-container">
          <div className="upload-box">
            <input type="file" accept='.jpg, .png' onChange={onChangeFile}/>
            {previewFileUrl && 
              <img src={previewFileUrl} className="preview-image"/>
            }
          </div>
          <button className="submit-button" onClick={onClickUpload} disabled={inProcess}>{inProcess ? '解析中...' : '送信する'}</button>
          <p className="note">※アップロードできるファイル形式は、JPEG/PNGのみです。</p>
      </div>
      <div className="result-container">
        <h3>解析結果</h3>
        {resultText && 
          <div style={{whiteSpace: "pre-line"}}>{resultText}</div>
        }
      </div>
    </div>
  );
}

export default App;
