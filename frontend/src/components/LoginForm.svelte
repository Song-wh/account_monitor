<script>
  import { createEventDispatcher } from 'svelte';
  
  const dispatch = createEventDispatcher();
  
  let virtualAccount = '';
  let authKey = '';
  let loading = false;
  let error = '';

  async function handleLogin() {
    if (!virtualAccount || !authKey) {
      error = '계좌번호와 키값을 모두 입력해주세요.';
      return;
    }

    loading = true;
    error = '';

    try {
      const response = await fetch('/api/auth', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify({
          virtual_account: virtualAccount,
          auth_key: authKey
        })
      });

      const result = await response.json();

      if (response.ok) {
        dispatch('login-success', { virtualAccount, authKey });
      } else {
        error = result.message || '인증에 실패했습니다.';
      }
    } catch (err) {
      error = '서버 연결에 실패했습니다.';
    } finally {
      loading = false;
    }
  }

  function handleKeyPress(event) {
    if (event.key === 'Enter') {
      handleLogin();
    }
  }
</script>

<div class="login-container">
  <div class="login-card">
    <div class="login-header">
      <h2>🔐 계좌 인증</h2>
      <p>가상계좌번호와 인증키를 입력해주세요</p>
    </div>

    <form on:submit|preventDefault={handleLogin}>
      <div class="form-group">
        <label for="virtualAccount">가상계좌번호</label>
        <input
          id="virtualAccount"
          type="text"
          bind:value={virtualAccount}
          placeholder="예: 123-456-7890"
          on:keypress={handleKeyPress}
          disabled={loading}
        />
      </div>

      <div class="form-group">
        <label for="authKey">인증키</label>
        <input
          id="authKey"
          type="password"
          bind:value={authKey}
          placeholder="인증키를 입력하세요"
          on:keypress={handleKeyPress}
          disabled={loading}
        />
      </div>

      {#if error}
        <div class="error-message">
          <span class="error-icon">⚠️</span>
          {error}
        </div>
      {/if}

      <button type="submit" class="login-button" disabled={loading}>
        {#if loading}
          <div class="button-spinner"></div>
          인증 중...
        {:else}
          🔑 인증하기
        {/if}
      </button>
    </form>
  </div>
</div>

<style>
  .login-container {
    width: 100%;
    display: flex;
    justify-content: center;
    align-items: center;
    min-height: 400px;
  }

  .login-card {
    background: white;
    border-radius: 16px;
    padding: 40px;
    box-shadow: 0 10px 30px rgba(0, 0, 0, 0.1);
    width: 100%;
    max-width: 400px;
  }

  .login-header {
    text-align: center;
    margin-bottom: 30px;
  }

  .login-header h2 {
    margin: 0 0 10px 0;
    color: #2c3e50;
    font-size: 1.8rem;
    font-weight: 600;
  }

  .login-header p {
    margin: 0;
    color: #6c757d;
    font-size: 0.95rem;
  }

  .form-group {
    margin-bottom: 20px;
  }

  .form-group label {
    display: block;
    margin-bottom: 8px;
    color: #495057;
    font-weight: 500;
    font-size: 0.9rem;
  }

  .form-group input {
    width: 100%;
    padding: 12px 16px;
    border: 2px solid #e9ecef;
    border-radius: 8px;
    font-size: 1rem;
    transition: all 0.3s ease;
    box-sizing: border-box;
  }

  .form-group input:focus {
    outline: none;
    border-color: #667eea;
    box-shadow: 0 0 0 3px rgba(102, 126, 234, 0.1);
  }

  .form-group input:disabled {
    background-color: #f8f9fa;
    cursor: not-allowed;
  }

  .error-message {
    background: #f8d7da;
    color: #721c24;
    padding: 12px 16px;
    border-radius: 8px;
    margin-bottom: 20px;
    display: flex;
    align-items: center;
    gap: 8px;
    font-size: 0.9rem;
  }

  .error-icon {
    font-size: 1.1rem;
  }

  .login-button {
    width: 100%;
    padding: 14px 20px;
    background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
    color: white;
    border: none;
    border-radius: 8px;
    font-size: 1rem;
    font-weight: 600;
    cursor: pointer;
    transition: all 0.3s ease;
    display: flex;
    align-items: center;
    justify-content: center;
    gap: 8px;
  }

  .login-button:hover:not(:disabled) {
    transform: translateY(-2px);
    box-shadow: 0 8px 20px rgba(102, 126, 234, 0.3);
  }

  .login-button:disabled {
    opacity: 0.7;
    cursor: not-allowed;
    transform: none;
  }

  .button-spinner {
    width: 16px;
    height: 16px;
    border: 2px solid rgba(255, 255, 255, 0.3);
    border-top: 2px solid white;
    border-radius: 50%;
    animation: spin 1s linear infinite;
  }

  @keyframes spin {
    0% { transform: rotate(0deg); }
    100% { transform: rotate(360deg); }
  }

  @media (max-width: 480px) {
    .login-card {
      padding: 30px 20px;
      margin: 0 10px;
    }
  }
</style>
