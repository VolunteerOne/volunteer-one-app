// This file was added by Matt
// It contains the component for a notification item card 
// which is dispayed on the noitifications screen accessible
// from the friends page. Used card component as a template.

import React from 'react';
import { withNavigation } from '@react-navigation/compat';
import PropTypes from 'prop-types';
import { StyleSheet, Dimensions, Image, TouchableWithoutFeedback } from 'react-native';
import { Block, Text, theme } from 'galio-framework';

import { argonTheme } from '../constants';


class NotificationItem extends React.Component {
  render() {
    const { navigation, item, horizontal, full, style, ctaColor, imageStyle } = this.props;
    
    const imageStyles = [
      full ? styles.fullImage : styles.horizontalImage,
      imageStyle
    ];
    const cardContainer = [styles.card, styles.shadow, style];
    const imgContainer = [styles.imageContainer,
      horizontal ? styles.horizontalStyles : styles.verticalStyles,
      styles.shadow
    ];

    return (
      <TouchableWithoutFeedback onPress={() => navigation.navigate('Pro')}> 
      <Block row={horizontal} card flex style={cardContainer}>
        



          <Block flex style={imgContainer}>
            <Image source={{uri: item.image}} 
            style={imageStyles} />
          </Block>

         

          <Block flex space="between" style={styles.cardDescription}>
            <Text size={14} style={styles.cardTitle}>{item.title}</Text>
            <Text size={12} muted={!ctaColor} color={ctaColor || argonTheme.COLORS.ACTIVE} bold>{item.time}h ago</Text>
          </Block>
        
      </Block>
      </TouchableWithoutFeedback>
    );
  }
}

NotificationItem.propTypes = {
  item: PropTypes.object,
  horizontal: PropTypes.bool,
  full: PropTypes.bool,
  ctaColor: PropTypes.string,
  imageStyle: PropTypes.any,
}

const styles = StyleSheet.create({
  card: {
    backgroundColor: theme.COLORS.WHITE,
    marginVertical: theme.SIZES.BASE,
    borderWidth: 0,
    minHeight: 114,
    marginBottom: 16
  },
  cardTitle: {
    flex: 1,
    flexWrap: 'wrap',
    paddingBottom: 6
  },
  cardDescription: {
    padding: theme.SIZES.BASE / 2
  },
  imageContainer: {
    borderRadius: 62,
    // elevation: 1,
    // overflow: 'hidden',
    marginTop: 17,
    marginBottom: 17,
    marginLeft: 17,
    marginRight: 10,
  },
  image: {
    // width: 124,
    // height: 124,
    // borderRadius: 62,
    // borderWidth: 0
  },
  horizontalImage: {
    height: 62,
    width: 62,
    borderRadius: 62,
  },
  horizontalStyles: {
    borderTopRightRadius: 62,
    borderBottomRightRadius: 62,
  },
  verticalStyles: {
    borderBottomRightRadius: 0,
    borderBottomLeftRadius: 0
  },
  fullImage: {
    height: 215
  },
  shadow: {
    shadowColor: theme.COLORS.BLACK,
    shadowOffset: { width: 0, height: 2 },
    shadowRadius: 4,
    shadowOpacity: 0.1,
    elevation: 2,
  },
});

export default withNavigation(NotificationItem);