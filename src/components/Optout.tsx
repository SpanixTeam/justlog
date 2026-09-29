import { IconButton, Button } from '@mui/material';
import { useContext, useState } from 'react';
import styled from 'styled-components';
import { store } from '../store';
import CancelIcon from '@mui/icons-material/Cancel';

const OptoutWrapper = styled.div``;

export function Optout() {
	const { state, setShowOptout } = useContext(store);

	const handleClick = () => {
		setShowOptout(!state.showOptout);
	};

	return (
		<OptoutWrapper>
			<IconButton
				aria-controls="docs"
				aria-haspopup="true"
				onClick={handleClick}
				size="small"
				color={state.showOptout ? 'primary' : 'default'}>
				<CancelIcon />
			</IconButton>
		</OptoutWrapper>
	);
}

const OptoutPanelWrapper = styled.div`
	background: var(--bg-bright);
	color: var(--text);
	margin: 3rem;
	font-size: 1.5rem;
	padding: 2rem;

	code {
		background: var(--bg);
		padding: 1rem;
		border-radius: 3px;
	}

	.generator {
		margin-top: 2rem;
		display: flex;
		gap: 1rem;
		align-items: center;

		input {
			background: var(--bg);
			border: none;
			color: white;
			padding: 0.6rem;
			font-size: 1.5rem;
			text-align: center;
			border-radius: 3px;
		}
	}

	.small {
		font-size: 0.8rem;
		font-family: monospace;
	}
`;

export function OptoutPanel() {
	const { state } = useContext(store);
	const [code, setCode] = useState('');

	const generateCode = () => {
		fetch(state.apiBaseUrl + '/optout', { method: 'POST' })
			.then((res) => res.json())
			.then(setCode)
			.catch(console.error);
	};

	return (
		<OptoutPanelWrapper>
			<p>
				Puedes desactivar el registro de tus mensajes. Esto también desactivará el 
				acceso a los datos registrados anteriormente.
				<br />
				Esto se aplica a todos los chats de esta instancia de Rustlog.
				<br />
				Darse de baja es definitivo, no se puede revertir. Así que piénsatelo bien 
				antes de darte de baja.
			</p>
			<p>
				Ten en cuenta que es posible que no aparezca ningún mensaje de confirmación.
			</p>
			<br />
			<div>
				<code>!rustlog optout {'<code>'}</code>
			</div>
			<div className="generator">
				<input readOnly type="text" value={code} />
				<Button variant="contained" onClick={generateCode} color="primary" size="large">
					Generar Código
				</Button>
			</div>
			{code && <p className="small">Este código es válido durante 60 segundos.</p>}
		</OptoutPanelWrapper>
	);
}
