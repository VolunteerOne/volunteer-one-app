// This file was added by Matt
// It contains the component for a profile item card 
// this item will be used when viewing a list of friends
// Used card component as a template.

import React from 'react';
import { withNavigation } from '@react-navigation/compat';
import PropTypes from 'prop-types';
import { StyleSheet, Dimensions, Image, TouchableWithoutFeedback } from 'react-native';
import { Block, Text, theme } from 'galio-framework';

import { argonTheme } from '../constants';
import { Button } from "../components";

class ProfileItem extends React.Component {
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
      <Block row={horizontal} card style={cardContainer}>

          {/* profile image */}
          <Block>
            <Image source={{uri: item.image}} style={imageStyles} />
          </Block>

          {/* username */}
          <Block flex style={styles.cardDescription}>
            <Text size={12} style={styles.cardTitle} bold>{item.username}</Text>
          </Block>

          {/* remove button */}
          <Block flex style={styles.cardDescription}>
            <Button
              small
              style={{ backgroundColor: 'grey' }}
            >
              Remove
            </Button>
          </Block>


      </Block>
      </TouchableWithoutFeedback>
    );
  }
}

ProfileItem.propTypes = {
  item: PropTypes.object,
  horizontal: PropTypes.bool,
  full: PropTypes.bool,
  ctaColor: PropTypes.string,
  imageStyle: PropTypes.any,
}

const styles = StyleSheet.create({
  card: {
    // backgroundColor: theme.COLORS.WHITE,
    // backgroundColor: 'blue',
    marginVertical: theme.SIZES.BASE /2,
    minHeight: 96,
    justifyContent: 'flex-start',
    alignItems: 'center',
  },
  cardTitle: {
    // flex: 1,
    // flexWrap: 'wrap',
    // paddingBottom: 10,
    // backgroundColor: 'red',
    // justifyContent: 'center',
    // alignItems: 'center', // matt
    // textAlignVertical: 'center',
  },
  cardDescription: {
    // paddingLeft: theme.SIZES.BASE / 5,  

    // paddingRight: theme.SIZES.BASE,  
    // paddingTop: theme.SIZES.BASE * 1.4,  
    // paddingBottom: theme.SIZES.BASE,  
    
    borderRadius: 3,
    elevation: 1,
    overflow: 'hidden',
    // backgroundColor: 'green'
  },
  image: {
  },
  horizontalImage: {
    height: 50,
    width: 50,
    borderRadius: 62,
    margin: 20,
  },
  horizontalStyles: {
    borderTopRightRadius: 0,
    borderBottomRightRadius: 0,
    
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

export default withNavigation(ProfileItem);