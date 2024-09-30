import { Stack, Tooltip } from '@mui/material';
import { ChatMessage } from '@twurple/chat';
import React, { useEffect, useRef, useState } from 'react';
import styled from 'styled-components';
import { ChannelBadge } from '../types/Badge';
import { useBadges } from '../hooks/useBadges';
import { useUserBadges } from '../hooks/useUserBadges';
import { CustomTooltip } from './Message';

const UserRoot = styled.div`
	margin-left: 5px;
	display: inline-flex;
	font-weight: 600;
	&:hover {
		text-shadow: 0 0 5px var(--theme3);
	}
`;

const UserContainer = styled.div.attrs((props) => ({
	style: {
		color: props.color,
	},
}))`
	display: inline;
	font-weight: 600;
	&:hover {
		text-shadow: 0 0 5px var(--theme3);
	}
`;

const UserBadge = styled.img`
	height: 1rem;
	display: inline;
	margin-right: 2px;
	border-radius: 0.15em;
	height: 1.1rem;
`;

export function User({
	displayName,
	color,
	badges,
	parsed,
}: {
	displayName: string;
	color: string;
	badges: string[];
	parsed: ChatMessage;
}): JSX.Element {
	const renderColor = color !== '' ? color : 'grey';
	// const [render, setRendering] = useState(false)
	// useEffect(() => {
	//   setRendering(true)
	// }, [])
	// if (!render) return null
	const channelBadges = useBadges(parsed.channelId);
	const userBadges = useUserBadges(parsed.userInfo?.userId);

	if (!parsed.id) return;

	return (
		<UserRoot>
			{badges.map((badgeId, i) => {
				const badge = channelBadges.get(badgeId);
				if (badge) {
					const url = badge.urls.small ?? null;
					if (!url) return null;
					return (
						<CustomTooltip
							key={`badge-${parsed.date.toISOString()}-${i}`}
							title={
								<Stack
									display="flex"
									alignItems="center"
									textAlign="center"
									gap="5px">
									<img
										src={badge.urls.big}
										style={{
											height: 72,
											width: 72,
											background: badge.color ?? 'transparent',
											borderRadius: '0.5em',
										}}
									/>
									<span>{badge.title}</span>
								</Stack>
							}>
							{badge.action ? (
								<a href={badge.action} target={badge.code}>
									<UserBadge
										src={url}
										style={{
											backgroundColor: badge.color ?? 'transparent',
										}}
									/>{' '}
								</a>
							) : (
								<UserBadge
									src={url}
									style={{
										backgroundColor: badge.color ?? 'transparent',
									}}
								/>
							)}
						</CustomTooltip>
					);
				}
				return null;
			})}
			{userBadges.badges.map((badge, i) => {
				const url = badge.urls.small ?? null;
				if (!url) return null;
				return (
					<CustomTooltip
						key={`badge-${parsed.date.toISOString()}-${i}`}
						title={
							<Stack display="flex" alignItems="center" textAlign="center" gap="5px">
								<img
									src={badge.urls.big}
									style={{
										height: 72,
										width: 72,
										background: badge.color ?? 'transparent',
										borderRadius: '0.5em',
									}}
								/>
								<span>{badge.title}</span>
							</Stack>
						}>
						{badge.action ? (
							<a href={badge.action} target={badge.code}>
								<UserBadge
									src={url}
									style={{ background: badge.color ?? 'transparent' }}
								/>{' '}
							</a>
						) : (
							<UserBadge
								src={url}
								style={{ background: badge.color ?? 'transparent' }}
							/>
						)}
					</CustomTooltip>
				);
			})}
			<UserContainer color={renderColor} className="user">
				{displayName}:
			</UserContainer>
		</UserRoot>
	);
}
